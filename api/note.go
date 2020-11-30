package kanka

import (
	"context"
	"fmt"
	"time"
)

// Notes is used to query the notes endpoints
type Notes struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Note is used to serialize a note object
type Note struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name     string `json:"name"`
	Entry    string `json:"entry"`
	NoteID   string `json:"note_id"`
	Type     string `json:"type"`
	IsPinned int    `json:"is_pinned"`

	Image          string `json:"image"`
	ImageFull      string `json:"image_full"`
	ImageThumb     string `json:"image_thumb"`
	HasCustomImage bool   `json:"has_custom_image"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// Notes returns a handle of the notes endpoint
func (c *Client) Notes(campaignID int) *Notes {
	return &Notes{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/notes", campaignID),
	}
}

// GetNotes can return information about all notes
func (n *Notes) GetNotes(ctx context.Context) (*[]Note, error) {

	var err error
	resp := []Note{}
	url := fmt.Sprintf("%s", n.urlPrefix)

	for len(url) > 0 && err == nil {
		page := []Note{}
		url, err = n.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetNote can return information about a single note
func (n *Notes) GetNote(ctx context.Context, id int) (*Note, error) {

	resp := Note{}
	_, err := n.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", n.urlPrefix, id), &resp)
	return &resp, err
}
