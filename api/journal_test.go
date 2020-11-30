package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJournals(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Journals(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/journals", c.urlPrefix)
}

func TestGetJournals(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	journals, err := client.Journals(1).GetJournals(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *journals, 1)

		// Main journal assertions
		j := (*journals)[0]
		assert.Equal(t, 1, j.ID)
		assert.Equal(t, 42, j.EntityID)
		assert.Equal(t, true, j.IsPrivate)
		assert.Equal(t, "Session 2 - Descent into the Abyss", j.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", j.Entry)
		assert.Equal(t, 11, j.CharacterID)
		assert.Equal(t, "2017-11-02", j.Date)
		assert.Equal(t, "Session", j.Type)
		assert.Equal(t, "https://example.com/image.png", j.Image)
		assert.Equal(t, "https://example.com/image_full.png", j.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", j.ImageThumb)
		assert.Equal(t, false, j.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, j.CreatedAt)
		assert.Equal(t, updated, j.UpdatedAt)
		assert.Equal(t, 1, j.CreatedBy)
		assert.Equal(t, 1, j.UpdatedBy)
	}
}

func TestGetJournal(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	journal, err := client.Journals(1).GetJournal(ctx, 1)

	if assert.NoError(t, err) {

		// Main journal assertions
		j := *journal
		assert.Equal(t, 1, j.ID)
		assert.Equal(t, 42, j.EntityID)
		assert.Equal(t, true, j.IsPrivate)
		assert.Equal(t, "Session 2 - Descent into the Abyss", j.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", j.Entry)
		assert.Equal(t, 11, j.CharacterID)
		assert.Equal(t, "2017-11-02", j.Date)
		assert.Equal(t, "Session", j.Type)
		assert.Equal(t, "https://example.com/image.png", j.Image)
		assert.Equal(t, "https://example.com/image_full.png", j.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", j.ImageThumb)
		assert.Equal(t, false, j.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, j.CreatedAt)
		assert.Equal(t, updated, j.UpdatedAt)
		assert.Equal(t, 1, j.CreatedBy)
		assert.Equal(t, 1, j.UpdatedBy)
	}
}
