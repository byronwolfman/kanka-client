package kanka

import (
	"context"
	"fmt"
	"time"
)

// Locations is used to query the locations endpoints
type Locations struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Location is used to serialize a location object
type Location struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name             string `json:"name"`
	Entry            string `json:"entry"`
	Type             string `json:"type"`
	Map              string `json:"map"`
	IsMapPrivate     int    `json:"is_map_private"`
	ParentLocationID int    `json:"parent_location_id"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// MapPoint is used to serialize a map point object
type MapPoint struct {
	Name           string `json:"name"`
	TargetEntityID int    `json:"target_entity_id"`

	AxisX  int    `json:"axis_x"`
	AxisY  int    `json:"axis_y"`
	Colour string `json:"colour"`
	Size   string `json:"size"`
	Icon   string `json:"icon"`
	Shape  string `json:"shape"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	addKnownLinkType("location", func(ctx context.Context, client *Client, campaignID int, entityID int) string {
		location, _ := client.Locations(campaignID).GetLocation(ctx, entityID)
		return location.Name
	})
}

// Locations returns a handle on the locations endpoints
func (c *Client) Locations(campaignID int) *Locations {
	return &Locations{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/locations", campaignID),
	}
}

// GetLocations can return information about all locations
func (l *Locations) GetLocations(ctx context.Context) (*[]Location, error) {

	var err error
	resp := []Location{}
	url := l.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Location{}
		url, err = l.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetLocation can return information about a single location
func (l *Locations) GetLocation(ctx context.Context, id int) (*Location, error) {

	resp := Location{}
	_, err := l.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", l.urlPrefix, id), &resp)
	return &resp, err
}

// GetMapPoints can return the map points of a given location
// Note that the API is documented as "map_points" but api docs claim it returns only a single item
func (l *Locations) GetMapPoints(ctx context.Context, id int) (*MapPoint, error) {

	resp := MapPoint{}
	_, err := l.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d/map_points", l.urlPrefix, id), &resp)
	return &resp, err
}
