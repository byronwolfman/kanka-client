package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAbilities(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Abilities(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/abilities", c.urlPrefix)
}

func TestGetAbilities(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	abilities, err := client.Abilities(1).GetAbilities(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *abilities, 1)

		// Main abilities assertions
		a := (*abilities)[0]
		assert.Equal(t, 1, a.ID)
		assert.Equal(t, 17, a.EntityID)
		assert.Equal(t, true, a.IsPrivate)
		assert.Equal(t, "Fireball", a.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", a.Entry)
		assert.Equal(t, 4, a.AbilityID)
		assert.Equal(t, 3, a.Charges)
		assert.Equal(t, "3rd level", a.Type)

		assert.Equal(t, "https://example.com/image.png", a.Image)
		assert.Equal(t, "https://example.com/image_full.png", a.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", a.ImageThumb)
		assert.Equal(t, false, a.HasCustomImage)

		// Date & time assertions
		created := time.Date(2020, time.March, 25, 13, 52, 42, 0, time.UTC)
		updated := time.Date(2020, time.May, 15, 8, 35, 56, 0, time.UTC)
		assert.Equal(t, created, a.CreatedAt)
		assert.Equal(t, updated, a.UpdatedAt)
		assert.Equal(t, 1, a.CreatedBy)
		assert.Equal(t, 1, a.UpdatedBy)
	}
}

func TestGetAbility(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	ability, err := client.Abilities(1).GetAbility(ctx, 1)

	if assert.NoError(t, err) {

		// Main abilities assertions
		a := *ability
		assert.Equal(t, 1, a.ID)
		assert.Equal(t, 17, a.EntityID)
		assert.Equal(t, true, a.IsPrivate)
		assert.Equal(t, "Fireball", a.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", a.Entry)
		assert.Equal(t, 4, a.AbilityID)
		assert.Equal(t, 3, a.Charges)
		assert.Equal(t, "3rd level", a.Type)

		assert.Equal(t, "https://example.com/image.png", a.Image)
		assert.Equal(t, "https://example.com/image_full.png", a.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", a.ImageThumb)
		assert.Equal(t, false, a.HasCustomImage)

		// Date & time assertions
		created := time.Date(2020, time.March, 25, 13, 52, 42, 0, time.UTC)
		updated := time.Date(2020, time.May, 15, 8, 35, 56, 0, time.UTC)
		assert.Equal(t, created, a.CreatedAt)
		assert.Equal(t, updated, a.UpdatedAt)
		assert.Equal(t, 1, a.CreatedBy)
		assert.Equal(t, 1, a.UpdatedBy)
	}
}
