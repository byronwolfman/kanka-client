package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRaces(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Races(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/races", c.urlPrefix)
}

func TestGetRaces(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	races, err := client.Races(1).GetRaces(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *races, 1)

		// Main race assertions
		r := (*races)[0]
		assert.Equal(t, 1, r.ID)
		assert.Equal(t, 7, r.EntityID)
		assert.Equal(t, true, r.IsPrivate)
		assert.Equal(t, "Goblin", r.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", r.Entry)
		assert.Equal(t, 3, r.RaceID)
		assert.Equal(t, "Goblinoid", r.Type)
		assert.Equal(t, "https://example.com/image.png", r.Image)
		assert.Equal(t, "https://example.com/image_full.png", r.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", r.ImageThumb)
		assert.Equal(t, false, r.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, r.CreatedAt)
		assert.Equal(t, updated, r.UpdatedAt)
		assert.Equal(t, 1, r.CreatedBy)
		assert.Equal(t, 1, r.UpdatedBy)
	}
}

func TestGetRace(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	race, err := client.Races(1).GetRace(ctx, 1)

	if assert.NoError(t, err) {

		// Main race assertions
		r := *race
		assert.Equal(t, 1, r.ID)
		assert.Equal(t, 7, r.EntityID)
		assert.Equal(t, true, r.IsPrivate)
		assert.Equal(t, "Goblin", r.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", r.Entry)
		assert.Equal(t, 3, r.RaceID)
		assert.Equal(t, "Goblinoid", r.Type)
		assert.Equal(t, "https://example.com/image.png", r.Image)
		assert.Equal(t, "https://example.com/image_full.png", r.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", r.ImageThumb)
		assert.Equal(t, false, r.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, r.CreatedAt)
		assert.Equal(t, updated, r.UpdatedAt)
		assert.Equal(t, 1, r.CreatedBy)
		assert.Equal(t, 1, r.UpdatedBy)
	}
}
