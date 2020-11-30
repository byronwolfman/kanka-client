package kanka

import (
	"context"
	"fmt"
	"time"
)

// Journals is used to query the journals endpoints
type Journals struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Journal is used to serialize an journal object
type Journal struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name        string `json:"name"`
	Entry       string `json:"entry"`
	CharacterID int    `json:"character_id"`
	Date        string `json:"date"`
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

// Journals returns a handle of the journals endpoint
func (c *Client) Journals(campaignID int) *Journals {
	return &Journals{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/journals", campaignID),
	}
}

// GetJournals can return information about all journals
func (j *Journals) GetJournals(ctx context.Context) (*[]Journal, error) {

	var err error
	resp := []Journal{}
	url := fmt.Sprintf("%s", j.urlPrefix)

	for len(url) > 0 && err == nil {
		page := []Journal{}
		url, err = j.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetJournal can return information about a single journal
func (j *Journals) GetJournal(ctx context.Context, id int) (*Journal, error) {

	resp := Journal{}
	_, err := j.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", j.urlPrefix, id), &resp)

	return &resp, err
}
