package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQuests(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Quests(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/quests", c.urlPrefix)
}

func TestGetQuests(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	quests, err := client.Quests(1).GetQuests(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *quests, 1)

		// Main quests assertions
		q := (*quests)[0]
		assert.Equal(t, 1, q.ID)
		assert.Equal(t, 164, q.EntityID)
		assert.Equal(t, true, q.IsPrivate)
		assert.Equal(t, "Pelor's Quest", q.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", q.Entry)
		assert.Equal(t, 4, q.CharacterID)
		assert.Equal(t, 2, q.Characters)
		assert.Equal(t, "2020-04-20", q.Date)
		assert.Equal(t, true, q.IsCompleted)
		assert.Equal(t, 1, q.Locations)
		assert.Equal(t, 3, q.QuestID)
		assert.Equal(t, "Main", q.Type)

		assert.Equal(t, "https://example.com/image.png", q.Image)
		assert.Equal(t, "https://example.com/image_full.png", q.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", q.ImageThumb)
		assert.Equal(t, false, q.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, q.CreatedAt)
		assert.Equal(t, updated, q.UpdatedAt)
		assert.Equal(t, 1, q.CreatedBy)
		assert.Equal(t, 1, q.UpdatedBy)
	}
}

func TestGetQuest(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	quest, err := client.Quests(1).GetQuest(ctx, 1)

	if assert.NoError(t, err) {

		// Main quests assertions
		q := *quest
		assert.Equal(t, 1, q.ID)
		assert.Equal(t, 164, q.EntityID)
		assert.Equal(t, true, q.IsPrivate)
		assert.Equal(t, "Pelor's Quest", q.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", q.Entry)
		assert.Equal(t, 4, q.CharacterID)
		assert.Equal(t, 2, q.Characters)
		assert.Equal(t, "2020-04-20", q.Date)
		assert.Equal(t, true, q.IsCompleted)
		assert.Equal(t, 1, q.Locations)
		assert.Equal(t, 3, q.QuestID)
		assert.Equal(t, "Main", q.Type)

		assert.Equal(t, "https://example.com/image.png", q.Image)
		assert.Equal(t, "https://example.com/image_full.png", q.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", q.ImageThumb)
		assert.Equal(t, false, q.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, q.CreatedAt)
		assert.Equal(t, updated, q.UpdatedAt)
		assert.Equal(t, 1, q.CreatedBy)
		assert.Equal(t, 1, q.UpdatedBy)
	}
}
