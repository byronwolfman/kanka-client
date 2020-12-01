package kanka

import (
	"context"
	"fmt"
	"time"
)

// Timelines is used to query the timelines endpoints
type Timelines struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Timeline is used to serialize an timeline object
type Timeline struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name        string `json:"name"`
	Entry       string `json:"entry"`
	Eras        []Era  `json:"eras"`
	RevertOrder bool   `json:"revert_order"`
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

// Era is used to serialize an era object
type Era struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	StartYear    int    `json:"start_year"`
	EndYear      int    `json:"end_year"`
}

// Timelines returns a handle of the timelines endpoint
func (c *Client) Timelines(campaignID int) *Timelines {
	return &Timelines{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/timelines", campaignID),
	}
}

// GetTimelines can return information about all timelines
func (t *Timelines) GetTimelines(ctx context.Context) (*[]Timeline, error) {

	var err error
	resp := []Timeline{}
	url := t.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Timeline{}
		url, err = t.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetTimeline can return information about a single timeline
func (t *Timelines) GetTimeline(ctx context.Context, id int) (*Timeline, error) {

	resp := Timeline{}
	_, err := t.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", t.urlPrefix, id), &resp)
	return &resp, err
}
