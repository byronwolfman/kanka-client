package kanka

import (
	"context"
	"fmt"
	"time"
)

// Campaigns is used to query the campaigns endpoints
type Campaigns struct {
	client *Client
}

// Campaign is used to serialize a campaign object
type Campaign struct {
	ID int `json:"id"`

	Name   string `json:"name"`
	Locale string `json:"locale"`
	Entry  string `json:"entry"`

	Image      string `json:"image"`
	ImageFull  string `json:"image_full"`
	ImageThumb string `json:"image_thumb"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Visibility string           `json:"visibility"`
	Members    []CampaignMember `json:"members"`
}

// CampaignMember is used to serialize campaign member objects
type CampaignMember struct {
	ID   int          `json:"id"`
	User CampaignUser `json:"user"`
}

// CampaignUser is used to serialize a campaign user object
type CampaignUser struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// Campaigns returns a handle on the campaigns endpoints
func (c *Client) Campaigns() *Campaigns {
	return &Campaigns{client: c}
}

// GetCampaigns can return information about all campaigns
func (c *Campaigns) GetCampaigns(ctx context.Context) (*[]Campaign, error) {

	var err error
	resp := []Campaign{}
	url := "/campaigns"

	for len(url) > 0 && err == nil {
		page := []Campaign{}
		url, err = c.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetCampaign returns information about a single campaign
func (c *Campaigns) GetCampaign(ctx context.Context, id int) (*Campaign, error) {

	resp := Campaign{}
	_, err := c.client.makeRequest(ctx, "GET", fmt.Sprintf("/campaigns/%d", id), &resp)
	return &resp, err
}
