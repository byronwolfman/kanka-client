package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEvents(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Events(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/events", c.urlPrefix)
}

func TestGetEvents(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	events, err := client.Events(1).GetEvents(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *events, 1)

		// Main events assertions
		e := (*events)[0]
		assert.Equal(t, 1, e.ID)
		assert.Equal(t, 7, e.EntityID)
		assert.Equal(t, true, e.IsPrivate)
		assert.Equal(t, "Battle of Hadish", e.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", e.Entry)
		assert.Equal(t, "44-3-16", e.Date)
		assert.Equal(t, "4", e.LocationID)
		assert.Equal(t, "Battle", e.Type)
		assert.Equal(t, "https://example.com/image.png", e.Image)
		assert.Equal(t, "https://example.com/image_full.png", e.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", e.ImageThumb)
		assert.Equal(t, false, e.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, e.CreatedAt)
		assert.Equal(t, updated, e.UpdatedAt)
		assert.Equal(t, 1, e.CreatedBy)
		assert.Equal(t, 1, e.UpdatedBy)
	}
}

func TestGetEvent(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	event, err := client.Events(1).GetEvent(ctx, 1)

	if assert.NoError(t, err) {

		// Main events assertions
		e := *event
		assert.Equal(t, 1, e.ID)
		assert.Equal(t, 7, e.EntityID)
		assert.Equal(t, true, e.IsPrivate)
		assert.Equal(t, "Battle of Hadish", e.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", e.Entry)
		assert.Equal(t, "44-3-16", e.Date)
		assert.Equal(t, "4", e.LocationID)
		assert.Equal(t, "Battle", e.Type)
		assert.Equal(t, "https://example.com/image.png", e.Image)
		assert.Equal(t, "https://example.com/image_full.png", e.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", e.ImageThumb)
		assert.Equal(t, false, e.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, e.CreatedAt)
		assert.Equal(t, updated, e.UpdatedAt)
		assert.Equal(t, 1, e.CreatedBy)
		assert.Equal(t, 1, e.UpdatedBy)
	}
}
