package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Tags(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/tags", c.urlPrefix)
}

func TestGetTags(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	tags, err := client.Tags(1).GetTags(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *tags, 1)

		// Main tag assertions
		tag := (*tags)[0]
		assert.Equal(t, 1, tag.ID)
		assert.Equal(t, 11, tag.EntityID)
		assert.Equal(t, true, tag.IsPrivate)
		assert.Equal(t, "Religion", tag.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", tag.Entry)
		assert.Equal(t, "green", tag.Colour)
		assert.Len(t, tag.Entities, 2)
		assert.Equal(t, 440, tag.Entities[1])
		assert.Equal(t, 4, tag.TagID)
		assert.Equal(t, "Lore", tag.Type)
		assert.Equal(t, "https://example.com/image.png", tag.Image)
		assert.Equal(t, "https://example.com/image_full.png", tag.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", tag.ImageThumb)
		assert.Equal(t, false, tag.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, tag.CreatedAt)
		assert.Equal(t, updated, tag.UpdatedAt)
		assert.Equal(t, 1, tag.CreatedBy)
		assert.Equal(t, 1, tag.UpdatedBy)
	}
}

func TestGetTag(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	tag, err := client.Tags(1).GetTag(ctx, 1)

	if assert.NoError(t, err) {

		// Main tag assertions
		assert.Equal(t, 1, tag.ID)
		assert.Equal(t, 11, tag.EntityID)
		assert.Equal(t, true, tag.IsPrivate)
		assert.Equal(t, "Religion", tag.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", tag.Entry)
		assert.Equal(t, "green", tag.Colour)
		assert.Len(t, tag.Entities, 2)
		assert.Equal(t, 440, tag.Entities[1])
		assert.Equal(t, 4, tag.TagID)
		assert.Equal(t, "Lore", tag.Type)
		assert.Equal(t, "https://example.com/image.png", tag.Image)
		assert.Equal(t, "https://example.com/image_full.png", tag.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", tag.ImageThumb)
		assert.Equal(t, false, tag.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, tag.CreatedAt)
		assert.Equal(t, updated, tag.UpdatedAt)
		assert.Equal(t, 1, tag.CreatedBy)
		assert.Equal(t, 1, tag.UpdatedBy)
	}
}
