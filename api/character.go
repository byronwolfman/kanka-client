package kanka

import (
	"context"
	"fmt"
	"time"
)

// Characters is used to query the characters endpoints
type Characters struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Character is used to serialize a character object
type Character struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Entry       string  `json:"entry"`
	EntryParsed string  `json:"entry_parsed"`
	LocationID  int     `json:"location_id"`
	Age         string  `json:"age"`
	Sex         string  `json:"sex"`
	RaceID      int     `json:"race_id"`
	Type        string  `json:"type"`
	FamilyID    int     `json:"family_id"`
	IsDead      bool    `json:"is_dead"`
	Traits      []Trait `json:"traits"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Trait is used to serialize a character trait object
type Trait struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Entry        string `json:"entry"`
	Section      string `json:"section"`
	IsPrivate    bool   `json:"is_private"`
	DefaultOrder int    `json:"default_order"`
}

// Characters returns a handle on the characters endpoint
func (c *Client) Characters(campaignID int) *Characters {
	return &Characters{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/characters", campaignID),
	}
}

// GetCharacters can return information about all characters
func (c *Characters) GetCharacters(ctx context.Context) (*[]Character, error) {

	var err error
	resp := []Character{}
	url := fmt.Sprintf("%s", c.urlPrefix)

	for len(url) > 0 && err == nil {
		fmt.Println("Getting", url)
		page := []Character{}
		url, err = c.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetCharacter can return information about a single character
func (c *Characters) GetCharacter(ctx context.Context, id int) (*Character, error) {

	resp := Character{}
	_, err := c.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", c.urlPrefix, id), &resp)
	return &resp, err
}
