package kanka

import (
	"context"
	"fmt"
	"time"
)

// Tags is used to query the tags endpoints
type Tags struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Tag is used to serialize an tag object
type Tag struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name     string `json:"name"`
	Entry    string `json:"entry"`
	Colour   string `json:"colour"`
	Entities []int  `json:"entities"`
	TagID    int    `json:"tag_id"`
	Type     string `json:"type"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Tags returns a handle of the tags endpoint
func (c *Client) Tags(campaignID int) *Tags {
	return &Tags{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/tags", campaignID),
	}
}

// GetTags can return information about all tags
func (t *Tags) GetTags(ctx context.Context) (*[]Tag, error) {

	var err error
	resp := []Tag{}
	url := fmt.Sprintf("%s", t.urlPrefix)

	for len(url) > 0 && err == nil {
		page := []Tag{}
		url, err = t.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetTag can return information about a single tag
func (t *Tags) GetTag(ctx context.Context, id int) (*Tag, error) {

	resp := Tag{}
	_, err := t.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", t.urlPrefix, id), &resp)

	return &resp, err
}
