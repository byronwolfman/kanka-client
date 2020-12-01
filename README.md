# Kanka Client

This is an unofficial API client library for the excellent worldbuilding site, [Kanka](https://kanka.io/). The methods in this library follow the API as closely as possible with certain caveats. In particular, most of the core object `GET` methods are covered, but not all, and even then not all attributes are serialized.

At the initial commit, the mocks have been built (and lightly modified) using the [1.0 API documentation](https://kanka.io/en-US/docs/1.0/) but not the live site, so some data may not be correctly serialized (but this will hopefully converge over time).

## Documentation

Generated automagically at https://godoc.org/github.com/byronwolfman/kanka-client/api.

## Example Usage

Note that error-checking is skipped for brevity. The API client itself can be setup with a few calls:

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

You can ask for a list of all objects of a given type. Be careful in that the client doesn't rate-limit itself in any way at the moment.
    
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

The default client configuration should be suitable for most cases, but can be modified if needed:

```go
client := kanka.NewClient(
	&kanka.Config{
		BaseURL:  "https://example.com/api/1.0", // Defaults to https://kanka.io/api/1.0
		ForceTLS: false,                         // Defaults to true
		Token:    "1234",                        // Defaults to ""
		Timeout:  time.Second * 30,              // Defaults to 15 seconds
	},
)
```

The `ForceTLS` parameter bears some explaining. When enabled, a config passed with a plain-HTTP base URL will be upgraded when the client initializes:

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

The reason for `ForceTLS` is to promote secure by the default, but also because sometimes Kanka will return a pagination object with a URL to the next page with a plain HTTP protocol. The extra measure behind `ForceTLS` makes a best effort to avoid sending such traffic in plaintext.

Probably the only reason to disable `ForceTLS` is for writing tests (which uses a localhost proxy) or to do traffic inspection through your own MitM proxy.

## Contributing

Pull requests that fix bugs and add test fixtures are welcome. Pull requests that add missing API endpoints or missing attributes for existing API endpoints are also very welcome.

## Further Reading

* [Kanka Website](http://kanka.io/)
* [Kanka 1.0 API Docs](https://kanka.io/en-US/docs/1.0)
* [Miscellany on Github](https://github.com/ilestis/miscellany) (the OSS powering Kanka)
