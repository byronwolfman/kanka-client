package kanka

import (
	"context"
	"fmt"
	"time"
)

// Races is used to query the races endpoints
type Races struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Race is used to serialize an race object
type Race struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name   string `json:"name"`
	Entry  string `json:"entry"`
	RaceID int    `json:"race_id"`
	Type   string `json:"type"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Races returns a handle of the races endpoint
func (c *Client) Races(campaignID int) *Races {
	return &Races{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/races", campaignID),
	}
}

// GetRaces can return information about all races
func (r *Races) GetRaces(ctx context.Context) (*[]Race, error) {

	var err error
	resp := []Race{}
	url := r.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Race{}
		url, err = r.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetRace can return information about a single race
func (r *Races) GetRace(ctx context.Context, id int) (*Race, error) {

	resp := Race{}
	_, err := r.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", r.urlPrefix, id), &resp)

	return &resp, err
}
