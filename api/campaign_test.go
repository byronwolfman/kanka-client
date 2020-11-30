package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCampaigns(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Campaigns()

	assert.Equal(t, client, c.client)
}

func TestGetCampaigns(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	campaigns, err := client.Campaigns().GetCampaigns(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *campaigns, 1)

		// Main campaign assertions
		c := (*campaigns)[0]
		assert.Equal(t, 1, c.ID)
		assert.Equal(t, "Thaelia", c.Name)
		assert.Equal(t, "en", c.Locale)
		assert.Equal(t, "\r\n<p>Aenean sit amet vehicula.</p>\r\n", c.Entry)
		assert.Equal(t, "https://example.com/image.png", c.Image)
		assert.Equal(t, "https://example.com/image_full.png", c.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", c.ImageThumb)
		assert.Equal(t, "public", c.Visibility)

		// Date & time assertions
		created := time.Date(2017, time.November, 2, 16, 29, 34, 0, time.UTC)
		updated := time.Date(2020, time.May, 23, 22, 0, 2, 0, time.UTC)
		assert.Equal(t, created, c.CreatedAt)
		assert.Equal(t, updated, c.UpdatedAt)

		// Membership assertions
		assert.Len(t, c.Members, 1)
		assert.Equal(t, 1, c.Members[0].ID)
		assert.Equal(t, 1, c.Members[0].User.ID)
		assert.Equal(t, "Ilestis", c.Members[0].User.Name)
		assert.Equal(t, "https://example.com/avatar.png", c.Members[0].User.Avatar)
	}
}

func TestGetCampaign(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	campaign, err := client.Campaigns().GetCampaign(ctx, 1)

	if assert.NoError(t, err) {

		// Main campaign assertions
		c := *campaign
		assert.Equal(t, 1, c.ID)
		assert.Equal(t, "Thaelia", c.Name)
		assert.Equal(t, "fr", c.Locale)
		assert.Equal(t, "\r\n<p>Aenean sit amet vehicula.</p>\r\n", c.Entry)
		assert.Equal(t, "https://example.com/image.png", c.Image)
		assert.Equal(t, "https://example.com/image_full.png", c.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", c.ImageThumb)
		assert.Equal(t, "public", c.Visibility)

		// Date & time assertions
		created := time.Date(2017, time.November, 2, 16, 29, 34, 0, time.UTC)
		updated := time.Date(2020, time.May, 23, 22, 0, 2, 0, time.UTC)
		assert.Equal(t, created, c.CreatedAt)
		assert.Equal(t, updated, c.UpdatedAt)

		// Membership assertions
		assert.Len(t, c.Members, 1)
		assert.Equal(t, 1, c.Members[0].ID)
		assert.Equal(t, 1, c.Members[0].User.ID)
		assert.Equal(t, "Ilestis", c.Members[0].User.Name)
		assert.Equal(t, "https://example.com/avatar.png", c.Members[0].User.Avatar)
	}
}
