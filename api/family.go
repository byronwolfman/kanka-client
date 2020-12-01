package kanka

import (
	"context"
	"fmt"
	"time"
)

// Families is used to query the families endpoints
type Families struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Family is used to serialize a family object
type Family struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name       string   `json:"name"`
	Entry      string   `json:"entry"`
	LocationID int      `json:"location_id"`
	FamilyID   int      `json:"family_id"`
	Type       string   `json:"type"`
	Members    []string `json:"members"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Families returns a handle on the families endpoint
func (c *Client) Families(campaignID int) *Families {
	return &Families{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/families", campaignID),
	}
}

// GetFamilies can return information about all families
func (f *Families) GetFamilies(ctx context.Context) (*[]Family, error) {

	var err error
	resp := []Family{}
	url := f.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Family{}
		url, err = f.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetFamily can return information about a single family
func (f *Families) GetFamily(ctx context.Context, id int) (*Family, error) {

	resp := Family{}
	_, err := f.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", f.urlPrefix, id), &resp)
	return &resp, err
}
