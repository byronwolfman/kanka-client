package kanka

import (
	"context"
	"fmt"
	"time"
)

// Items is used to query the items endpoints
type Items struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Item is used to serialize an item object
type Item struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name        string `json:"name"`
	Entry       string `json:"entry"`
	LocationID  int    `json:"location_id"`
	CharacterID int    `json:"character_id"`
	Price       string `json:"price"`
	Size        string `json:"size"`
	Type        string `json:"type"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Items returns a handle of the items endpoint
func (c *Client) Items(campaignID int) *Items {
	return &Items{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/items", campaignID),
	}
}

// GetItems can return information about all items
func (i *Items) GetItems(ctx context.Context) (*[]Item, error) {

	var err error
	resp := []Item{}
	url := i.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Item{}
		url, err = i.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetItem can return information about a single item
func (i *Items) GetItem(ctx context.Context, id int) (*Item, error) {

	resp := Item{}
	_, err := i.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", i.urlPrefix, id), &resp)

	return &resp, err
}
