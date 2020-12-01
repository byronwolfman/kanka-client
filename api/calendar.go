package kanka

import (
	"context"
	"fmt"
	"time"
)

// Calendars is used to query the calendars endpoints
type Calendars struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Calendar is used to serialize an calendar object
type Calendar struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name  string `json:"name"`
	Entry string `json:"entry"`
	Date  string `json:"date"`
	Type  string `json:"type"`

	Weekdays []string          `json:"weekdays"`
	Months   []Month           `json:"months"`
	Years    map[string]string `json:"years"`

	Moons       []Moon   `json:"moons"`
	Seasons     []Season `json:"seasons"`
	StartOffset int      `json:"start_offset"`

	Suffix         string `json:"suffix"`
	HasLeapYear    bool   `json:"has_leap_year"`
	LeapYearAmount int    `json:"leap_year_amount"`
	LeapYearMonth  int    `json:"leap_year_month"`
	LeapYearOffset int    `json:"leap_year_offset"`
	LeapYearStart  int    `json:"leap_year_start"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Month is used to serialize a month object
type Month struct {
	Name   string `json:"name"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

// Moon is used to serialize a moon object
type Moon struct {
	Name     string `json:"name"`
	Fullmoon string `json:"fullmoon"`
	Offset   int    `json:"offset"`
	Colour   string `json:"colour"`
}

// Season is used to serialize a season object
type Season struct {
	Name  string `json:"name"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
}

// Calendars returns a handle of the calendars endpoint
func (c *Client) Calendars(campaignID int) *Calendars {
	return &Calendars{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/calendars", campaignID),
	}
}

// GetCalendars can return information about all calendars
func (c *Calendars) GetCalendars(ctx context.Context) (*[]Calendar, error) {

	var err error
	resp := []Calendar{}
	url := c.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Calendar{}
		url, err = c.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetCalendar can return information about a single calendar
func (c *Calendars) GetCalendar(ctx context.Context, id int) (*Calendar, error) {

	resp := Calendar{}
	_, err := c.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", c.urlPrefix, id), &resp)
	return &resp, err
}
