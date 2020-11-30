package kanka

import (
	"context"
	"fmt"
	"time"
)

// Abilities is used to query the abilities endpoints
type Abilities struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Ability is used to serialize an ability object
type Ability struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name      string `json:"name"`
	Entry     string `json:"entry"`
	AbilityID int    `json:"ability_id"`
	Charges   int    `json:"charges"`
	Type      string `json:"type"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Abilities returns a handle of the abilities endpoint
func (c *Client) Abilities(campaignID int) *Abilities {
	return &Abilities{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/abilities", campaignID),
	}
}

// GetAbilities can return information about all abilities
func (a *Abilities) GetAbilities(ctx context.Context) (*[]Ability, error) {

	var err error
	resp := []Ability{}
	url := fmt.Sprintf("%s", a.urlPrefix)

	for len(url) > 0 && err == nil {
		page := []Ability{}
		url, err = a.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetAbility can return information about a single ability
func (a *Abilities) GetAbility(ctx context.Context, id int) (*Ability, error) {

	resp := Ability{}
	_, err := a.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", a.urlPrefix, id), &resp)

	return &resp, err
}
