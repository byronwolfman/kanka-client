package kanka

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {

	config := DefaultConfig()

	assert.Equal(t, BaseURLV1, config.BaseURL)
	assert.True(t, config.ForceTLS)
	assert.Equal(t, time.Second*15, config.Timeout)
	assert.Empty(t, config.Token)
}

func TestNewClient(t *testing.T) {

	// Verify that client matches passed configs
	client := NewClient(DefaultConfig())

	assert.Equal(t, BaseURLV1, client.BaseURL)
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

func TestMakeRequest(t *testing.T) {

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

		res.Write([]byte("body"))
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

// mockTestServer is a convenience function to return an httptest server
// that loads a generic mock HTTP 200 JSON response.
func mockTestServer() (*httptest.Server, *Config) {

	// Create the server
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		// Mock file will be at a filepath that matches the URL path
		mockFile := fmt.Sprintf("mocks/%s/%s.json", req.Method, req.URL.Path)
		body, err := ioutil.ReadFile(mockFile)
		if err != nil {
			res.WriteHeader(404)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write([]byte(body))
	}))

	// Create a client config for reaching the server
	config := DefaultConfig()
	config.BaseURL = testServer.URL
	config.ForceTLS = false

	return testServer, config
}
