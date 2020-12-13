/*
Package kanka implements a basic client library for the Kanka API.
See https://kanka.io/docs/1.0 for API documentation.
*/
package kanka

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	// KankaBaseURLV1 is the default 1.0 base URL for the Kanka API
	KankaBaseURLV1 = "https://kanka.io/api/1.0"
	// DefaultMaxRequestsPerMinute is the number of requests per minute before client self-rate-limits
	DefaultMaxRequestsPerMinute = 30
	// DefaultTimeout for http client request timeouts
	DefaultTimeout = time.Second * 15
)

// Client provides a client to the Kanka API
type Client struct {
	BaseURL                string
	ForceTLS               bool
	rateLimiter            chan int
	rateLimitResetInterval time.Duration
	token                  string
	HTTPClient             *http.Client
}

// Config is used to configure the creation of a client
type Config struct {
	BaseURL              string
	ForceTLS             bool
	MaxRequestsPerMinute time.Duration
	Token                string
	Timeout              time.Duration
}

// Response is used to serialize a successful API response
type Response struct {
	Data  interface{} `json:"data"`
	Links Links       `json:"links,omitempty"`
	Meta  Meta        `json:"meta,omitempty"`
}

// Links is used to serialize paginated responses
type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}

// Meta is used to serialize metadata of responses
type Meta struct {
	Path        string `json:"path"`
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	PerPage     int    `json:"per_page"`
	From        int    `json:"from"`
	To          int    `json:"to"`
	Total       int    `json:"total"`
}

// DefaultConfig returns a default configuration for the client
func DefaultConfig() *Config {
	return &Config{
		BaseURL:              KankaBaseURLV1,
		MaxRequestsPerMinute: DefaultMaxRequestsPerMinute,
		Timeout:              DefaultTimeout,
		ForceTLS:             true,
	}
}

// NewClient returns a new client
func NewClient(c *Config) *Client {

	// Assert https if ForceTLS is true
	if strings.HasPrefix(c.BaseURL, "http://") && c.ForceTLS {
		c.BaseURL = strings.Replace(c.BaseURL, "http", "https", 1)
	}

	// Assert that config's MaxRequestsPerMinute is above zero
	if c.MaxRequestsPerMinute <= 0 {
		c.MaxRequestsPerMinute = DefaultMaxRequestsPerMinute
	}

	// Create and return client
	return &Client{
		BaseURL:                c.BaseURL,
		ForceTLS:               c.ForceTLS,
		rateLimiter:            make(chan int, c.MaxRequestsPerMinute),
		rateLimitResetInterval: time.Minute,
		token:                  c.Token,
		HTTPClient: &http.Client{
			Timeout: c.Timeout,
		},
	}
}

// SetRateLimitResetInterval adjusts how quickly requests are unblocked by the rate-limiter channel.
// This should almost never be modified. The rate-limiter receives MaxRequestsPerMinute from a Config
// object which assumes that the reset interval is 1 minute. The rate-limiter will work incorrectly
// if the reset interval is changed to anything other than 1 minute. The only reason to change the
// rate-limit interval is for convenience when writing tests for the rate-limiter so that it's not
// necessary to wait a full minute for tests to pass. Outside of this use case, avoid modifying it.
//
// Another more dangerous way to think about this function is that it effectively adjusts how long
// 1 minute is considered by the rate-limiter.
func (c *Client) SetRateLimitResetInterval(duration time.Duration) {
	c.rateLimitResetInterval = duration
}

// makeRequest is a convenience function to make a request at a given URL endpoint, and then decode the response into v
// Populates the response into interface v
// Returns URL of the next page (if one exists) and an error (if one exists)
func (c *Client) makeRequest(ctx context.Context, method string, endpoint string, v interface{}) (string, error) {

	// Sometimes endpoint is just the path, e.g. /campaigns
	// Sometimes, if we're paginating, it will be the full URL, e.g. https://example.com/campaigns

	// For the latter case, double-check TLS and upgrade if necessary
	if strings.HasPrefix(endpoint, "http://") && c.ForceTLS {
		endpoint = strings.Replace(endpoint, "http", "https", 1)
	}

	// Also for the latter case: make sure the base URL hasn't been modified, and then strip it
	if strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://") {
		if !strings.HasPrefix(endpoint, c.BaseURL) {
			return "", fmt.Errorf("Base URL in request to '%s' does not match %s", endpoint, c.BaseURL)
		}
		endpoint = strings.Replace(endpoint, c.BaseURL, "", 1)
	}

	// Also recheck the base URL in case it's been modified
	if strings.HasPrefix(c.BaseURL, "http://") && c.ForceTLS {
		c.BaseURL = strings.Replace(c.BaseURL, "http", "https", 1)
	}

	// Setup the request
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, endpoint), nil)
	if err != nil {
		return "", err
	}

	// Add headers
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	// Self-rate limit; the channel will block if we request too fast
	c.rateLimiter <- 1
	defer func() {
		go func() {
			time.Sleep(c.rateLimitResetInterval)
			<-c.rateLimiter
		}()
	}()

	// Make the reqest
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// Bail out non-2xx responses (http.Client.Do follows redirects)
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("Non-2xx response: %s", resp.Status)
	}

	// Bail out for non-JSON responses
	respContentHeader := resp.Header.Get("Content-Type")
	if !strings.Contains(strings.ToLower(respContentHeader), "application/json") {
		return "", fmt.Errorf("Non-JSON response: %s, Content-Type: %s", resp.Status, respContentHeader)
	}

	// Attempt to decode the response into a Response object (Data takes the shape of v)
	fullResponse := Response{
		Data: v,
	}

	// Return err on failed JSON decode
	if err = json.NewDecoder(resp.Body).Decode(&fullResponse); err != nil {
		return "", err
	}

	// If there is another page, return the URL to it
	if fullResponse.Meta.CurrentPage < fullResponse.Meta.LastPage {
		return fullResponse.Links.Next, nil
	}

	// Otherwise return no next page, no error
	return "", nil
}
