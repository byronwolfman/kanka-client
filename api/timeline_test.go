package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimelines(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Timelines(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/timelines", c.urlPrefix)
}

func TestGetTimelines(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	timelines, err := client.Timelines(1).GetTimelines(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *timelines, 1)

		// Main timelines assertions
		tl := (*timelines)[0]
		assert.Equal(t, 1, tl.ID)
		assert.Equal(t, 112, tl.EntityID)
		assert.Equal(t, true, tl.IsPrivate)
		assert.Equal(t, "Thaelian Timeline", tl.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", tl.Entry)
		assert.Equal(t, false, tl.RevertOrder)
		assert.Equal(t, "Primary", tl.Type)
		assert.Equal(t, "https://example.com/image.png", tl.Image)
		assert.Equal(t, "https://example.com/image_full.png", tl.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", tl.ImageThumb)
		assert.Equal(t, false, tl.HasCustomImage)

		// Eras
		assert.Len(t, tl.Eras, 2)
		assert.Equal(t, "Before Common Era", tl.Eras[1].Name)
		assert.Equal(t, "BCE", tl.Eras[1].Abbreviation)
		assert.Equal(t, 0, tl.Eras[1].StartYear)
		assert.Equal(t, 0, tl.Eras[1].EndYear)

		// Date & time assertions
		created := time.Date(2019, time.January, 28, 6, 29, 29, 0, time.UTC)
		updated := time.Date(2020, time.January, 30, 17, 30, 52, 0, time.UTC)
		assert.Equal(t, created, tl.CreatedAt)
		assert.Equal(t, updated, tl.UpdatedAt)
		assert.Equal(t, 1, tl.CreatedBy)
		assert.Equal(t, 1, tl.UpdatedBy)
	}
}

func TestGetTimeline(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	timeline, err := client.Timelines(1).GetTimeline(ctx, 1)

	if assert.NoError(t, err) {

		// Main timelineassertions
		tl := *timeline
		assert.Equal(t, 1, tl.ID)
		assert.Equal(t, 112, tl.EntityID)
		assert.Equal(t, true, tl.IsPrivate)
		assert.Equal(t, "Thaelian Timeline", tl.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", tl.Entry)
		assert.Equal(t, false, tl.RevertOrder)
		assert.Equal(t, "Primary", tl.Type)
		assert.Equal(t, "https://example.com/image.png", tl.Image)
		assert.Equal(t, "https://example.com/image_full.png", tl.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", tl.ImageThumb)
		assert.Equal(t, false, tl.HasCustomImage)

		// Eras
		assert.Len(t, tl.Eras, 2)
		assert.Equal(t, "Before Common Era", tl.Eras[1].Name)
		assert.Equal(t, "BCE", tl.Eras[1].Abbreviation)
		assert.Equal(t, 0, tl.Eras[1].StartYear)
		assert.Equal(t, 0, tl.Eras[1].EndYear)

		// Date & time assertions
		created := time.Date(2019, time.January, 28, 6, 29, 29, 0, time.UTC)
		updated := time.Date(2020, time.January, 30, 17, 30, 52, 0, time.UTC)
		assert.Equal(t, created, tl.CreatedAt)
		assert.Equal(t, updated, tl.UpdatedAt)
		assert.Equal(t, 1, tl.CreatedBy)
		assert.Equal(t, 1, tl.UpdatedBy)
	}
}
