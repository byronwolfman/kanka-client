package kanka

import (
	"context"
	"fmt"
	"time"
)

// Events is used to query the events endpoints
type Events struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Event is used to serialize an event object
type Event struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name       string `json:"name"`
	Entry      string `json:"entry"`
	Date       string `json:"date"`
	LocationID string `json:"location_id"`
	Type       string `json:"type"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Events returns a handle of the events endpoint
func (c *Client) Events(campaignID int) *Events {
	return &Events{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/events", campaignID),
	}
}

// GetEvents can return information about all events
func (e *Events) GetEvents(ctx context.Context) (*[]Event, error) {

	var err error
	resp := []Event{}
	url := e.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Event{}
		url, err = e.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetEvent can return information about a single event
func (e *Events) GetEvent(ctx context.Context, id int) (*Event, error) {

	resp := Event{}
	_, err := e.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", e.urlPrefix, id), &resp)
	return &resp, err
}
