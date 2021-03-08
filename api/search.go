package kanka

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

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
