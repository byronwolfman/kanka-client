package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCharacters(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Characters(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/characters", c.urlPrefix)
}

func TestGetCharacters(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	characters, err := client.Characters(1).GetCharacters(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *characters, 1)

		// Main character assertions
		c := (*characters)[0]
		assert.Equal(t, 1, c.ID)
		assert.Equal(t, 4, c.EntityID)
		assert.Equal(t, true, c.IsPrivate)
		assert.Equal(t, "Jonathan Green", c.Name)
		assert.Equal(t, "The Hero", c.Title)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", c.Entry)
		assert.Equal(t, 4, c.LocationID)
		assert.Equal(t, "39", c.Age)
		assert.Equal(t, "Male", c.Sex)
		assert.Equal(t, 3, c.RaceID)
		assert.Equal(t, "Player Character", c.Type)
		assert.Equal(t, 34, c.FamilyID)
		assert.Equal(t, true, c.IsDead)
		assert.Equal(t, "https://example.com/image.png", c.Image)
		assert.Equal(t, "https://example.com/image_full.png", c.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", c.ImageThumb)
		assert.Equal(t, false, c.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 29, 16, 40, 34, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 38, 46, 0, time.UTC)
		assert.Equal(t, created, c.CreatedAt)
		assert.Equal(t, updated, c.UpdatedAt)
		assert.Equal(t, 1, c.CreatedBy)
		assert.Equal(t, 1, c.UpdatedBy)

		// Traits
		assert.Len(t, c.Traits, 1)

		tr := c.Traits[0]
		assert.Equal(t, 33, tr.ID)
		assert.Equal(t, "Goals", tr.Name)
		assert.Equal(t, "Become a Paladin.", tr.Entry)
		assert.Equal(t, "personality", tr.Section)
		assert.Equal(t, false, tr.IsPrivate)
		assert.Equal(t, 0, tr.DefaultOrder)
	}
}

func TestGetCharacter(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	character, err := client.Characters(1).GetCharacter(ctx, 1)

	if assert.NoError(t, err) {

		// Main character assertions
		c := *character
		assert.Equal(t, 1, c.ID)
		assert.Equal(t, 4, c.EntityID)
		assert.Equal(t, true, c.IsPrivate)
		assert.Equal(t, "Jonathan Green", c.Name)
		assert.Equal(t, "The Hero", c.Title)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", c.Entry)
		assert.Equal(t, 4, c.LocationID)
		assert.Equal(t, "39", c.Age)
		assert.Equal(t, "Male", c.Sex)
		assert.Equal(t, 3, c.RaceID)
		assert.Equal(t, "Player Character", c.Type)
		assert.Equal(t, 34, c.FamilyID)
		assert.Equal(t, true, c.IsDead)
		assert.Equal(t, "https://example.com/image.png", c.Image)
		assert.Equal(t, "https://example.com/image_full.png", c.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", c.ImageThumb)
		assert.Equal(t, false, c.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 29, 16, 40, 34, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 38, 46, 0, time.UTC)
		assert.Equal(t, created, c.CreatedAt)
		assert.Equal(t, updated, c.UpdatedAt)
		assert.Equal(t, 1, c.CreatedBy)
		assert.Equal(t, 1, c.UpdatedBy)

		// Traits
		tr := c.Traits[0]
		assert.Equal(t, 33, tr.ID)
		assert.Equal(t, "Goals", tr.Name)
		assert.Equal(t, "Become a Paladin.", tr.Entry)
		assert.Equal(t, "personality", tr.Section)
		assert.Equal(t, false, tr.IsPrivate)
		assert.Equal(t, 0, tr.DefaultOrder)
	}
}
