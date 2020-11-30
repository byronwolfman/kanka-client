package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotes(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Notes(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/notes", c.urlPrefix)
}

func TestGetNotes(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	notes, err := client.Notes(1).GetNotes(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *notes, 1)

		// Main notes assertions
		n := (*notes)[0]
		assert.Equal(t, 2, n.ID)
		assert.Equal(t, 7, n.EntityID)
		assert.Equal(t, true, n.IsPrivate)
		assert.Equal(t, "Legends of the World", n.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", n.Entry)
		assert.Equal(t, "1", n.NoteID)
		assert.Equal(t, "Lore", n.Type)
		assert.Equal(t, 0, n.IsPinned)
		assert.Equal(t, "https://example.com/image.png", n.Image)
		assert.Equal(t, "https://example.com/image_full.png", n.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", n.ImageThumb)
		assert.Equal(t, false, n.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, n.CreatedAt)
		assert.Equal(t, updated, n.UpdatedAt)
		assert.Equal(t, 1, n.CreatedBy)
		assert.Equal(t, 1, n.UpdatedBy)
	}
}

func TestGetNote(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	note, err := client.Notes(1).GetNote(ctx, 1)

	if assert.NoError(t, err) {

		// Main notes assertions
		n := *note
		assert.Equal(t, 2, n.ID)
		assert.Equal(t, 7, n.EntityID)
		assert.Equal(t, true, n.IsPrivate)
		assert.Equal(t, "Legends of the World", n.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", n.Entry)
		assert.Equal(t, "1", n.NoteID)
		assert.Equal(t, "Lore", n.Type)
		assert.Equal(t, 0, n.IsPinned)
		assert.Equal(t, "https://example.com/image.png", n.Image)
		assert.Equal(t, "https://example.com/image_full.png", n.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", n.ImageThumb)
		assert.Equal(t, false, n.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, n.CreatedAt)
		assert.Equal(t, updated, n.UpdatedAt)
		assert.Equal(t, 1, n.CreatedBy)
		assert.Equal(t, 1, n.UpdatedBy)
	}
}
