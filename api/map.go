package kanka

import (
	"context"
	"fmt"
	"time"
)

// Maps is used to query the maps endpoints
type Maps struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Map is used to serialize an map object
type Map struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name        string `json:"name"`
	Entry       string `json:"entry"`
	EntryParsed string `json:"entry_parsed"`
	LocationID  int    `json:"location_id"`
	MapID       int    `json:"map_id"`
	Type        string `json:"type"`

	Grid    int `json:"grid"`
	CenterX int `json:"center_x"`
	CenterY int `json:"center_y"`
	Height  int `json:"height"`
	Width   int `json:"width"`

	InitialZoom int `json:"initial_zoom"`
	MinZoom     int `json:"min_zoom"`
	MaxZoom     int `json:"max_zoom"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// MapMarker is used to serialize an map marker object
type MapMarker struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name       string `json:"name"`
	MapID      int    `json:"map_id"`
	SizeID     int    `json:"size_id"`
	Visibility string `json:"visibility"`

	Colour     string `json:"colour"`
	FontColour string `json:"font_colour"`
	Opacity    int    `json:"opacity"`

	CustomIcon string `json:"custom_icon"`
	Icon       string `json:"icon"`

	CustomShape string `json:"custom_shape"`
	ShapeID     int    `json:"shape_id"`

	IsDraggable bool   `json:"is_draggable"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// MapGroup is used to serialize an map group object
type MapGroup struct {
	ID        int  `json:"id"`
	IsPrivate bool `json:"is_private"`

	Name       string `json:"name"`
	IsShown    bool   `json:"is_shown"`
	MapID      int    `json:"map_id"`
	Position   int    `json:"position"`
	Visibility string `json:"visibility"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Maps returns a handle of the maps endpoint
func (c *Client) Maps(campaignID int) *Maps {
	return &Maps{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/maps", campaignID),
	}
}

// GetMaps can return information about all maps
func (m *Maps) GetMaps(ctx context.Context) (*[]Map, error) {

	var err error
	resp := []Map{}
	url := m.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Map{}
		url, err = m.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetMap can return information about a single map
func (m *Maps) GetMap(ctx context.Context, id int) (*Map, error) {

	resp := Map{}
	_, err := m.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", m.urlPrefix, id), &resp)

	return &resp, err
}

// GetMapMarkers can return information about all map markers for a given map
func (m *Maps) GetMapMarkers(ctx context.Context, id int) (*[]MapMarker, error) {

	var err error
	resp := []MapMarker{}
	url := fmt.Sprintf("%s/%d/map_markers", m.urlPrefix, id)

	for len(url) > 0 && err == nil {
		page := []MapMarker{}
		url, err = m.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetMapGroups can return information about all map groups for a given map
func (m *Maps) GetMapGroups(ctx context.Context, id int) (*[]MapGroup, error) {

	var err error
	resp := []MapGroup{}
	url := fmt.Sprintf("%s/%d/map_groups", m.urlPrefix, id)

	for len(url) > 0 && err == nil {
		page := []MapGroup{}
		url, err = m.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}
