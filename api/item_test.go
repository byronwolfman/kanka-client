package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestItems(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Items(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/items", c.urlPrefix)
}

func TestGetItems(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	items, err := client.Items(1).GetItems(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *items, 1)

		// Main item assertions
		i := (*items)[0]
		assert.Equal(t, 1, i.ID)
		assert.Equal(t, 7, i.EntityID)
		assert.Equal(t, true, i.IsPrivate)
		assert.Equal(t, "Spear", i.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", i.Entry)
		assert.Equal(t, 4, i.LocationID)
		assert.Equal(t, 2, i.CharacterID)
		assert.Equal(t, "25 gp", i.Price)
		assert.Equal(t, "1 lb.", i.Size)
		assert.Equal(t, "Weapon", i.Type)
		assert.Equal(t, "https://example.com/image.png", i.Image)
		assert.Equal(t, "https://example.com/image_full.png", i.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", i.ImageThumb)
		assert.Equal(t, false, i.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, i.CreatedAt)
		assert.Equal(t, updated, i.UpdatedAt)
		assert.Equal(t, 1, i.CreatedBy)
		assert.Equal(t, 1, i.UpdatedBy)
	}
}

func TestGetItem(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	item, err := client.Items(1).GetItem(ctx, 1)

	if assert.NoError(t, err) {

		// Main item assertions
		i := *item
		assert.Equal(t, 1, i.ID)
		assert.Equal(t, 7, i.EntityID)
		assert.Equal(t, true, i.IsPrivate)
		assert.Equal(t, "Spear", i.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", i.Entry)
		assert.Equal(t, 4, i.LocationID)
		assert.Equal(t, 2, i.CharacterID)
		assert.Equal(t, "25 gp", i.Price)
		assert.Equal(t, "1 lb.", i.Size)
		assert.Equal(t, "Weapon", i.Type)
		assert.Equal(t, "https://example.com/image.png", i.Image)
		assert.Equal(t, "https://example.com/image_full.png", i.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", i.ImageThumb)
		assert.Equal(t, false, i.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, i.CreatedAt)
		assert.Equal(t, updated, i.UpdatedAt)
		assert.Equal(t, 1, i.CreatedBy)
		assert.Equal(t, 1, i.UpdatedBy)
	}
}
