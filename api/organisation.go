package kanka

import (
	"context"
	"fmt"
	"time"
)

// Organisations is used to query the organisations endpoints
type Organisations struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Organisation is used to serialize an organisation object
type Organisation struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name           string `json:"name"`
	Entry          string `json:"entry"`
	LocationID     int    `json:"location_id"`
	OrganisationID int    `json:"organisation_id"`
	Type           string `json:"type"`
	Members        int    `json:"members"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

func init() {
	addKnownLinkType("organisation", func(ctx context.Context, client *Client, campaignID int, entityID int) string {
		organisation, _ := client.Organisations(campaignID).GetOrganisation(ctx, entityID)
		return organisation.Name
	})
}

// Organisations returns a handle of the organisations endpoint
func (c *Client) Organisations(campaignID int) *Organisations {
	return &Organisations{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/organisations", campaignID),
	}
}

// GetOrganisations can return information about all organisations
func (o *Organisations) GetOrganisations(ctx context.Context) (*[]Organisation, error) {

	var err error
	resp := []Organisation{}
	url := o.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Organisation{}
		url, err = o.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetOrganisation can return information about a single organisation
func (o *Organisations) GetOrganisation(ctx context.Context, id int) (*Organisation, error) {

	resp := Organisation{}
	_, err := o.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", o.urlPrefix, id), &resp)
	return &resp, err
}
