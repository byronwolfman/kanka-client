package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrganisations(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Organisations(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/organisations", c.urlPrefix)
}

func TestGetOrganisations(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	organisations, err := client.Organisations(1).GetOrganisations(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *organisations, 1)

		// Main organisation assertions
		o := (*organisations)[0]
		assert.Equal(t, 1, o.ID)
		assert.Equal(t, 5, o.EntityID)
		assert.Equal(t, true, o.IsPrivate)
		assert.Equal(t, "Tiamat Cultists", o.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", o.Entry)
		assert.Equal(t, 4, o.LocationID)
		assert.Equal(t, 4, o.OrganisationID)
		assert.Equal(t, "Kingdom", o.Type)
		assert.Equal(t, 3, o.Members)
		assert.Equal(t, "https://example.com/image.png", o.Image)
		assert.Equal(t, "https://example.com/image_full.png", o.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", o.ImageThumb)
		assert.Equal(t, false, o.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, o.CreatedAt)
		assert.Equal(t, updated, o.UpdatedAt)
		assert.Equal(t, 1, o.CreatedBy)
		assert.Equal(t, 1, o.UpdatedBy)

	}
}

func TestGetOrganisation(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	organisation, err := client.Organisations(1).GetOrganisation(ctx, 1)

	if assert.NoError(t, err) {

		// Main organisation assertions
		o := *organisation
		assert.Equal(t, 1, o.ID)
		assert.Equal(t, 5, o.EntityID)
		assert.Equal(t, true, o.IsPrivate)
		assert.Equal(t, "Tiamat Cultists", o.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", o.Entry)
		assert.Equal(t, 4, o.LocationID)
		assert.Equal(t, 4, o.OrganisationID)
		assert.Equal(t, "Kingdom", o.Type)
		assert.Equal(t, 3, o.Members)
		assert.Equal(t, "https://example.com/image.png", o.Image)
		assert.Equal(t, "https://example.com/image_full.png", o.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", o.ImageThumb)
		assert.Equal(t, false, o.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 0, 1, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, o.CreatedAt)
		assert.Equal(t, updated, o.UpdatedAt)
		assert.Equal(t, 1, o.CreatedBy)
		assert.Equal(t, 1, o.UpdatedBy)

	}
}
