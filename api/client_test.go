package kanka

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {

	config := DefaultConfig()

	assert.Equal(t, KankaBaseURLV1, config.BaseURL)
	assert.True(t, config.ForceTLS)
	assert.Equal(t, time.Second*15, config.Timeout)
	assert.Empty(t, config.Token)
}

func TestNewClient(t *testing.T) {

	// Verify that client matches passed configs
	client := NewClient(DefaultConfig())

	assert.Equal(t, KankaBaseURLV1, client.BaseURL)
	assert.True(t, client.ForceTLS)
	assert.Equal(t, time.Second*15, client.HTTPClient.Timeout)
	assert.Empty(t, client.token)

	// Make sure that protocol is ugpraded to https if ForceTLS is true
	client = NewClient(
		&Config{
			BaseURL:  "http://example.com/api/1.0",
			ForceTLS: true,
		},
	)

	assert.Equal(t, "https://example.com/api/1.0", client.BaseURL)
}

func TestRequestResponses(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		switch path := req.URL.Path; path {
		case "/404":
			res.WriteHeader(404)
		case "/bad-content-type":
			res.Header().Set("Content-Type", "text/html")
			res.WriteHeader(200)
		default:
			res.Header().Set("Content-Type", "fail")
			res.WriteHeader(200)
		}

		_, err := res.Write([]byte("body"))
		assert.NoError(t, err)
	}))
	defer func() { testServer.Close() }()

	// Create client
	config := DefaultConfig()
	config.BaseURL = testServer.URL
	config.ForceTLS = false

	client := NewClient(config)
	ctx := context.Background()

	// Test 404s
	_, err := client.makeRequest(ctx, "GET", "/404", nil)
	assert.EqualError(t, err, "Non-2xx response: 404 Not Found")

	// Test Content-Type header
	_, err = client.makeRequest(ctx, "GET", "/bad-content-type", nil)
	assert.EqualError(t, err, "Non-JSON response: 200 OK, Content-Type: text/html")
}

func TestSelfRateLimit(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		_, err := res.Write([]byte("{}"))
		assert.NoError(t, err)
	}))
	defer func() { testServer.Close() }()

	// Create client
	config := DefaultConfig()
	config.BaseURL = testServer.URL
	config.ForceTLS = false
	config.MaxRequestsPerMinute = 30

	client := NewClient(config)
	client.SetRateLimitResetInterval(2 * time.Second) // Speed up tests
	ctx := context.Background()

	// 30 requests should complete quickly with no rate-limiting; 100ms ought to be more than enough
	start := time.Now()
	for i := 0; i < 30; i++ {
		_, err := client.makeRequest(ctx, "GET", "/", nil)
		assert.NoError(t, err)
	}
	end := time.Now()
	assert.True(t, end.Before(start.Add(100*time.Millisecond)))

	// Create new client to reset rate-limit
	client = NewClient(config)
	client.SetRateLimitResetInterval(2 * time.Second) // Speed up tests

	// 31 requests should complete with ~ 2 seconds of rate-limiting
	start = time.Now()
	for i := 0; i < 31; i++ {
		_, err := client.makeRequest(ctx, "GET", "/", nil)
		assert.NoError(t, err)
	}
	end = time.Now()
	assert.True(t, end.Before(start.Add(2500*time.Millisecond)))
	assert.True(t, end.After(start.Add(2000*time.Millisecond)))
}

// mockTestServer is a convenience function to return an httptest server
// that loads a generic mock HTTP 200 JSON response.
func mockTestServer() (*httptest.Server, *Config) {

	// Create the server
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		// Mock file will be at a filepath that matches the URL path
		mockFile := fmt.Sprintf("mocks/%s/%s.json", req.Method, req.URL.Path)
		body, err := ioutil.ReadFile(filepath.Clean(mockFile))
		if err != nil {
			res.WriteHeader(404)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		_, err = res.Write([]byte(body))

		if err != nil {
			panic(fmt.Sprintf("mockTestServer encountered error while writing response: %v", err))
		}
	}))

	// Create a client config for reaching the server
	config := DefaultConfig()
	config.BaseURL = testServer.URL
	config.ForceTLS = false

	return testServer, config
}
