package kanka

import (
	"context"
	"fmt"
	"time"
)

// Quests is used to query the quests endpoints
type Quests struct {
	client     *Client
	campaignID int
	urlPrefix  string
}

// Quest is used to serialize an quest object
type Quest struct {
	ID        int  `json:"id"`
	EntityID  int  `json:"entity_id"`
	IsPrivate bool `json:"is_private"`

	Name        string `json:"name"`
	Entry       string `json:"entry"`
	CharacterID int    `json:"character_id"`
	Characters  int    `json:"characters"`
	Date        string `json:"date"`
	IsCompleted bool   `json:"is_completed"`
	Locations   int    `json:"locations"`
	QuestID     int    `json:"quest_id"`
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

// QuestCharacter is used to serialize a quest character object
type QuestCharacter struct {
	ID          int    `json:"id"`
	CharacterID int    `json:"character_id"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// QuestItem is used to serialize a quest item object
type QuestItem struct {
	ID          int    `json:"id"`
	ItemID      int    `json:"character_id"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// QuestLocation is used to serialize a quest location object
type QuestLocation struct {
	ID          int    `json:"id"`
	LocationID  int    `json:"location_id"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// QuestOrganisation is used to serialize a quest organization object
type QuestOrganisation struct {
	ID             int    `json:"id"`
	OrganisationID int    `json:"location_id"`
	Description    string `json:"description"`
	IsPrivate      bool   `json:"is_private"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

func init() {
	addKnownLinkType("quest", func(ctx context.Context, client *Client, campaignID int, entityID int) string {
		quest, _ := client.Quests(campaignID).GetQuest(ctx, entityID)
		return quest.Name
	})
}

// Quests returns a handle of the quests endpoint
func (c *Client) Quests(campaignID int) *Quests {
	return &Quests{
		client:     c,
		campaignID: campaignID,
		urlPrefix:  fmt.Sprintf("/campaigns/%d/quests", campaignID),
	}
}

// GetQuests can return information about all quests
func (q *Quests) GetQuests(ctx context.Context) (*[]Quest, error) {

	var err error
	resp := []Quest{}
	url := q.urlPrefix

	for len(url) > 0 && err == nil {
		page := []Quest{}
		url, err = q.client.makeRequest(ctx, "GET", url, &page)
		resp = append(resp, page...)
	}

	return &resp, err
}

// GetQuest can return information about a single quest
func (q *Quests) GetQuest(ctx context.Context, id int) (*Quest, error) {

	resp := Quest{}
	_, err := q.client.makeRequest(ctx, "GET", fmt.Sprintf("%s/%d", q.urlPrefix, id), &resp)
	return &resp, err
}
