package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLocations(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Locations(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/locations", c.urlPrefix)
}

func TestGetLocations(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	locations, err := client.Locations(1).GetLocations(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *locations, 1)

		// Main location assertions
		l := (*locations)[0]
		assert.Equal(t, 1, l.ID)
		assert.Equal(t, 5, l.EntityID)
		assert.Equal(t, true, l.IsPrivate)
		assert.Equal(t, "Mordor", l.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", l.Entry)
		assert.Equal(t, "Kingdom", l.Type)
		assert.Equal(t, "https://example.com/map", l.Map)
		assert.Equal(t, 0, l.IsMapPrivate)
		assert.Equal(t, 4, l.ParentLocationID)
		assert.Equal(t, "https://example.com/image.png", l.Image)
		assert.Equal(t, "https://example.com/image_full.png", l.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", l.ImageThumb)
		assert.Equal(t, false, l.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, l.CreatedAt)
		assert.Equal(t, updated, l.UpdatedAt)
		assert.Equal(t, 1, l.CreatedBy)
		assert.Equal(t, 1, l.UpdatedBy)
	}
}

func TestGetLocation(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	location, err := client.Locations(1).GetLocation(ctx, 1)

	if assert.NoError(t, err) {

		// Main location assertions
		l := *location
		assert.Equal(t, 1, l.ID)
		assert.Equal(t, 5, l.EntityID)
		assert.Equal(t, true, l.IsPrivate)
		assert.Equal(t, "Mordor", l.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", l.Entry)
		assert.Equal(t, "Kingdom", l.Type)
		assert.Equal(t, "https://example.com/map", l.Map)
		assert.Equal(t, 0, l.IsMapPrivate)
		assert.Equal(t, 4, l.ParentLocationID)
		assert.Equal(t, "https://example.com/image.png", l.Image)
		assert.Equal(t, "https://example.com/image_full.png", l.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", l.ImageThumb)
		assert.Equal(t, false, l.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, l.CreatedAt)
		assert.Equal(t, updated, l.UpdatedAt)
		assert.Equal(t, 1, l.CreatedBy)
		assert.Equal(t, 1, l.UpdatedBy)
	}
}

func TestGetMapPoints(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	m, err := client.Locations(1).GetMapPoints(ctx, 1)

	if assert.NoError(t, err) {

		// Main map points assertions
		assert.Equal(t, 58, m.TargetEntityID)
		assert.Equal(t, "Somewhere", m.Name)
		assert.Equal(t, 1356, m.AxisX)
		assert.Equal(t, 788, m.AxisY)
		assert.Equal(t, "red", m.Colour)
		assert.Equal(t, "small", m.Size)
		assert.Equal(t, "skull", m.Icon)
		assert.Equal(t, "circle", m.Shape)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, m.CreatedAt)
		assert.Equal(t, updated, m.UpdatedAt)
	}
}
