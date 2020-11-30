package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMaps(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Maps(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/maps", c.urlPrefix)
}

func TestGetMaps(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	maps, err := client.Maps(1).GetMaps(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *maps, 1)

		// Main map assertions
		m := (*maps)[0]
		assert.Equal(t, 1, m.ID)
		assert.Equal(t, 164, m.EntityID)
		assert.Equal(t, true, m.IsPrivate)
		assert.Equal(t, "Pelor's Map", m.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", m.Entry)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", m.EntryParsed)
		assert.Equal(t, 4, m.LocationID)
		assert.Equal(t, 5, m.MapID)
		assert.Equal(t, "Continent", m.Type)
		assert.Equal(t, "https://example.com/image.png", m.Image)
		assert.Equal(t, "https://example.com/image_full.png", m.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", m.ImageThumb)
		assert.Equal(t, false, m.HasCustomImage)

		assert.Equal(t, 0, m.Grid)
		assert.Equal(t, 1, m.CenterX)
		assert.Equal(t, 1, m.CenterY)
		assert.Equal(t, 1080, m.Height)
		assert.Equal(t, 1920, m.Width)
		assert.Equal(t, -1, m.InitialZoom)
		assert.Equal(t, -1, m.MinZoom)
		assert.Equal(t, 10, m.MaxZoom)

		// Date & time assertions
		created := time.Date(2020, time.September, 18, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2020, time.September, 18, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, m.CreatedAt)
		assert.Equal(t, updated, m.UpdatedAt)
		assert.Equal(t, 1, m.CreatedBy)
		assert.Equal(t, 1, m.UpdatedBy)
	}
}

func TestGetMap(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	m, err := client.Maps(1).GetMap(ctx, 1)

	if assert.NoError(t, err) {

		// Main map assertions
		assert.Equal(t, 1, m.ID)
		assert.Equal(t, 164, m.EntityID)
		assert.Equal(t, true, m.IsPrivate)
		assert.Equal(t, "Pelor's Map", m.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", m.Entry)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", m.EntryParsed)
		assert.Equal(t, 4, m.LocationID)
		assert.Equal(t, 5, m.MapID)
		assert.Equal(t, "Continent", m.Type)
		assert.Equal(t, "https://example.com/image.png", m.Image)
		assert.Equal(t, "https://example.com/image_full.png", m.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", m.ImageThumb)
		assert.Equal(t, false, m.HasCustomImage)

		assert.Equal(t, 0, m.Grid)
		assert.Equal(t, 1, m.CenterX)
		assert.Equal(t, 1, m.CenterY)
		assert.Equal(t, 1080, m.Height)
		assert.Equal(t, 1920, m.Width)
		assert.Equal(t, -1, m.InitialZoom)
		assert.Equal(t, -1, m.MinZoom)
		assert.Equal(t, 10, m.MaxZoom)

		// Date & time assertions
		created := time.Date(2020, time.September, 18, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2020, time.September, 18, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, m.CreatedAt)
		assert.Equal(t, updated, m.UpdatedAt)
		assert.Equal(t, 1, m.CreatedBy)
		assert.Equal(t, 1, m.UpdatedBy)
	}
}

func TestGetMapMarkers(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	markers, err := client.Maps(1).GetMapMarkers(ctx, 1)

	if assert.NoError(t, err) {
		assert.Len(t, *markers, 1)

		// Main map marker assertions
		m := (*markers)[0]
		assert.Equal(t, 31, m.ID)
		assert.Equal(t, 5, m.EntityID)
		assert.Equal(t, true, m.IsPrivate)
		assert.Equal(t, "Shape", m.Name)
		assert.Equal(t, 1, m.MapID)
		assert.Equal(t, 1, m.SizeID)
		assert.Equal(t, "all", m.Visibility)

		assert.Equal(t, "/en-US/docs/1.0/map_markers#008000", m.Colour)
		assert.Equal(t, "/en-US/docs/1.0/map_markers#000000", m.FontColour)
		assert.Equal(t, 100, m.Opacity)

		assert.Equal(t, "https://example.com/marker.png", m.CustomIcon)
		assert.Equal(t, "1", m.Icon)

		assert.Equal(t, "500,500 500,600, 600,600 600,500", m.CustomShape)
		assert.Equal(t, 5, m.ShapeID)

		assert.Equal(t, true, m.IsDraggable)
		assert.Equal(t, "422.857", m.Latitude)
		assert.Equal(t, "499.000", m.Longitude)

		// Date & time assertions
		created := time.Date(2020, time.July, 25, 10, 10, 30, 0, time.UTC)
		updated := time.Date(2020, time.July, 25, 10, 10, 35, 0, time.UTC)
		assert.Equal(t, created, m.CreatedAt)
		assert.Equal(t, updated, m.UpdatedAt)
		assert.Equal(t, 1, m.CreatedBy)
		assert.Equal(t, 1, m.UpdatedBy)
	}
}

func TestGetMapGroups(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	groups, err := client.Maps(1).GetMapGroups(ctx, 1)

	if assert.NoError(t, err) {
		assert.Len(t, *groups, 1)

		// Main map group assertions
		g := (*groups)[0]
		assert.Equal(t, 3, g.ID)
		assert.Equal(t, true, g.IsPrivate)
		assert.Equal(t, "Spoon", g.Name)
		assert.Equal(t, true, g.IsShown)
		assert.Equal(t, 1, g.MapID)
		assert.Equal(t, 1, g.Position)
		assert.Equal(t, "all", g.Visibility)

		// Date & time assertions
		created := time.Date(2020, time.July, 25, 16, 24, 34, 0, time.UTC)
		updated := time.Date(2020, time.July, 25, 16, 24, 39, 0, time.UTC)
		assert.Equal(t, created, g.CreatedAt)
		assert.Equal(t, updated, g.UpdatedAt)
		assert.Equal(t, 1, g.CreatedBy)
		assert.Equal(t, 1, g.UpdatedBy)
	}
}
