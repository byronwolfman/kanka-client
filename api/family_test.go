package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFamilies(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Families(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/families", c.urlPrefix)
}

func TestGetFamilies(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	families, err := client.Families(1).GetFamilies(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *families, 1)

		// Main family assertions
		f := (*families)[0]
		assert.Equal(t, 1, f.ID)
		assert.Equal(t, 5, f.EntityID)
		assert.Equal(t, true, f.IsPrivate)
		assert.Equal(t, "Adams", f.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", f.Entry)
		assert.Equal(t, 4, f.LocationID)
		assert.Equal(t, 2, f.FamilyID)
		assert.Equal(t, "Gothic", f.Type)
		assert.Len(t, f.Members, 1)
		assert.Equal(t, "3", f.Members[0])

		assert.Equal(t, "https://example.com/image.png", f.Image)
		assert.Equal(t, "https://example.com/image_full.png", f.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", f.ImageThumb)
		assert.Equal(t, false, f.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, f.CreatedAt)
		assert.Equal(t, updated, f.UpdatedAt)
		assert.Equal(t, 1, f.CreatedBy)
		assert.Equal(t, 1, f.UpdatedBy)
	}
}

func TestGetFamily(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	family, err := client.Families(1).GetFamily(ctx, 1)

	if assert.NoError(t, err) {

		// Main family assertions
		f := *family
		assert.Equal(t, 1, f.ID)
		assert.Equal(t, 5, f.EntityID)
		assert.Equal(t, true, f.IsPrivate)
		assert.Equal(t, "Adams", f.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", f.Entry)
		assert.Equal(t, 1, f.LocationID)
		assert.Equal(t, 2, f.FamilyID)
		assert.Equal(t, "Gothic", f.Type)
		assert.Len(t, f.Members, 1)
		assert.Equal(t, "3", f.Members[0])

		assert.Equal(t, "https://example.com/image.png", f.Image)
		assert.Equal(t, "https://example.com/image_full.png", f.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", f.ImageThumb)
		assert.Equal(t, false, f.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, f.CreatedAt)
		assert.Equal(t, updated, f.UpdatedAt)
		assert.Equal(t, 1, f.CreatedBy)
		assert.Equal(t, 1, f.UpdatedBy)
	}
}
