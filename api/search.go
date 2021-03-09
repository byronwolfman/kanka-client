package kanka

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var linkRe *regexp.Regexp = regexp.MustCompile(`\[[a-z]+:[0-9]+\]`)
var linkAlpha *regexp.Regexp = regexp.MustCompile(`[[:alpha:]]+`)
var linkDigit *regexp.Regexp = regexp.MustCompile(`[[:digit:]]+`)

var knownLinkTypes map[string]interface{} = make(map[string]interface{})

// Searches is used to query the search endpoint
type Searches struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// SearchResult is used to serialize responses from the search endpoint
type SearchResult struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name    string `json:"name"`
	ToolTip string `json:"tooltip"`
	Type    string `json:"type"`
	URL     string `json:"url"`

	Image          string `json:"image"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Searches returns a handle on the search endpoint
func (c *Client) Searches(campaignID int) *Searches {
	return &Searches{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/search", campaignID),
	}
}

// Search returns one or more results on a search request
func (s *Searches) Search(ctx context.Context, searchTerm string) (*[]SearchResult, error) {

	// Need to URL encode spaces, etc; url.QueryEscape produces `+` for spaces which appears to be incompatible with kanka
	encodedTerm := &url.URL{Path: searchTerm}

	var err error
	resp := []SearchResult{}
	url := fmt.Sprintf("%s/%s", s.urlPrefix, encodedTerm.String())

	for len(url) > 0 && err == nil {
		page := []SearchResult{}
		url, err = s.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// ResolveLinksToText attempts to lookup and resolve the names of linked entities in a string.
// E.g. attempt to resolve "Hailing from the city of [location:1234]" to "Hailing from the city of Neverwinter"
func (s *Searches) ResolveLinksToText(ctx context.Context, rawString string) string {

	// Extract [alpha:digit] instances from string
	links := linkRe.FindAllString(rawString, -1)
	for _, link := range links {

		// Extract the [type:ID]
		linkType := linkAlpha.FindString(link)
		linkIDAsString := linkDigit.FindString(link)

		// Quick check to make sure we didn't mess this up
		if linkType == "" || linkIDAsString == "" {
			continue
		}

		// Convert link ID to integer
		linkID, err := strconv.Atoi(linkIDAsString)
		if err != nil {
			continue
		}

		// Do we recognize this link type? Extract the function if so (otherwise, skip)
		linkFn, ok := knownLinkTypes[linkType]
		if !ok {
			continue
		}

		// Resolve the name (will be "" if resolution fails)
		resolvedName := linkFn.(func(context.Context, *Client, int, int) string)(ctx, s.client, s.campaignID, linkID)

		// Empty name means we probably failed to resolve it
		if resolvedName == "" {
			continue
		}

		// Rewrite the raw string
		rawString = strings.Replace(rawString, link, resolvedName, 1)
	}

	return rawString
}

// addKnownLinkType is used to register known link types (characters, locations) and the function to look up their name.
// It should be called by all entity source files (e.g. character.go) in their init function
func addKnownLinkType(typeName string, fn interface{}) {
	knownLinkTypes[typeName] = fn
}
