[![Go Reference](https://pkg.go.dev/badge/github.com/byronwolfman/kanka-client.svg)](https://pkg.go.dev/github.com/byronwolfman/kanka-client) [![Go Report Card](https://goreportcard.com/badge/github.com/byronwolfman/kanka-client)](https://goreportcard.com/report/github.com/byronwolfman/kanka-client)

# Kanka Client

This is an unofficial API client library for the excellent worldbuilding site, [Kanka](https://kanka.io/). The methods in this library follow the API as closely as possible with certain caveats. In particular, most of the core object `GET` methods are covered, but not all, and even then not all attributes are serialized.

At the initial commit, the mocks have been built (and lightly modified) using the [1.0 API documentation](https://kanka.io/en-US/docs/1.0/) but not the live site, so some data may not be correctly serialized (but this will hopefully converge over time).

## Example Usage

Note that error-checking is skipped for brevity.

### Client Setup

The API client itself can be setup with a few calls:

```go
package main

import (
	"context"
	"fmt"

	kanka "github.com/byronwolfman/kanka-client/api"
)

func main() {

	config := kanka.DefaultConfig()
	config.Token = "personal-access-token"

	client := kanka.NewClient(config)
	ctx := context.Background()
	...
}
```

The default client configuration should be suitable for most cases, but can be modified if needed:

```go
client := kanka.NewClient(
	&kanka.Config{
		BaseURL:              "https://example.com/api/1.0", // Defaults to https://kanka.io/api/1.0
		ForceTLS:             false,                         // Defaults to true
		MaxRequestsPerMinute: 90,                            // Defaults to 30
		Token:                "1234",                        // Defaults to ""
		Timeout:              time.Second * 30,              // Defaults to 15 seconds
	},
)
```

### Requesting Objects

You can ask for a list of all objects of a given type. Kanka returns paginated results of up to 15 objects per page which the client automatically depaginates. For particularly large responses this may run afoul of throttling, so the client makes a best effort to rate-limit itself (described in more detail further down).
    
```go
// Set the campaign ID
campaignID := 1234

// Get all characters in the campaign
characters, err := client.Characters(campaignID).GetCharacters(ctx)
for _, character := range *characters {
	fmt.Printf("%d: %s\n", character.ID, character.Name)
}
    
// Get all locations in the campaign
locations, err := client.Locations(campaignID).GetLocations(ctx)
for _, location := range *locations {
	fmt.Printf("%d: %s\n", location.ID, location.Name)
}
```

You can also ask for a specific entities:

```go
// Set the campaign ID
campaignID := 1234

// Get a specific location
location, err := client.Locations(campaignID).GetLocation(ctx, 56789)
fmt.Println(location.Entry)

// Get a specific item
item, err := client.Items(campaignID).GetItem(ctx, 23456)
fmt.Println(item.Name)
fmt.Println(item.Type)
```

### TLS Configuration

The `ForceTLS` parameter is enabled by default and bears some explaining. When enabled, a config passed with a plain-HTTP base URL will be upgraded when the client initializes:

```go
client := kanka.NewClient(
	&kanka.Config{
		BaseURL:  "http://example.com/api/1.0",
		ForceTLS: true,
	},
)

fmt.Println(client.BaseURL)

// Prints "https://example.com/api/1.0"
```

The protocol is re-checked on all API calls since the base URL could otherwise be changed back again with `client.BaseURL = http://example.com`.

The `ForceTLS` option exists to promote security by the default, but also because Kanka may return a pagination object with plaintext URLs which cannot be queried.

```sh
curl \
  -Ss \
  -H "Authorization: Bearer ${TOKEN}" \
  -H "Content-Type: application/json" \
  "https://kanka.io/api/1.0/campaigns/1234/characters" \
  | jq

{
  "data": [
    {...}
  ],
  "links": {
    "first": "http://kanka.io/api/1.0/campaigns/1234/characters?page=1",  # Plain HTTP
    "last": "http://kanka.io/api/1.0/campaigns/1234/characters?page=2",   # Plain HTTP
    "prev": null,
    "next": "http://kanka.io/api/1.0/campaigns/1234/characters?page=2"    # Plain HTTP
  },
  "meta": {...}
}
```

The API client is designed to follow these links in order to retrieve the paginated results. To avoid leaking the API token, the client makes two additional checks:

1. If `ForceTLS` is enabled but the link protocol is plain HTTP, then the link will be re-written as HTTPS before it is followed.
1. If the link's base URL does not match the client's configured base URL (including after being upgraded if `ForceTLS` is enabled) then the client will throw an error and refuse to follow the link.

Probably the only reason to disable `ForceTLS` is for writing tests (which uses a localhost proxy) or to do traffic inspection through your own MitM proxy.

### Rate-Limiting

The Kanka API allows a maximum of 30 requests per minute per client, or 90 requests per minute for subscribers. The Kanka API returns HTTP 429 if this limit is exceeded.

To make a best effort to avoid being throttled, the client rate-limits itself to 30 requests per minute by default:

```go
// Rate-limited to the default 30 requests per minute
client := kanka.NewClient(kanka.DefaultConfig())

// Also rate-limited to the default 30 requests per minute
client := kanka.NewClient(
	&kanka.Config{
		BaseURL: "https://example.com/api/1.0",
		Token:   "1234",
	},
)

// Rate-limited to 90 requests per minute, high enough for subscribers
client := kanka.NewClient(
	&kanka.Config{
		BaseURL:              "https://example.com/api/1.0",
		MaxRequestsPerMinute: 90,
		Token:                "1234",
	},
)
```

The limit is immutable at client creation time and therefore must be passed through a Config object. With the default limit, the client will be able to dispatch up to 30 requests in quick succession without slowing itself down. If the client needs to make a 31st request however, then that request will be blocked for 1 minute beginning from the moment the 1st request received its response.

## Contributing

Pull requests that fix bugs and add test fixtures are welcome. Pull requests that add missing API endpoints or missing attributes for existing API endpoints are also very welcome.

## Further Reading

* [Kanka Website](http://kanka.io/)
* [Kanka 1.0 API Docs](https://kanka.io/en-US/docs/1.0)
* [Miscellany on Github](https://github.com/ilestis/miscellany) (the OSS powering Kanka)
