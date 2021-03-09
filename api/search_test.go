package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSearches(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Searches(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/search", c.urlPrefix)
}

func TestSearch(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	results, err := client.Searches(1).Search(ctx, "tyr")

	if assert.NoError(t, err) {
		assert.Len(t, *results, 1)

		// Main search result assertions
		r := (*results)[0]
		assert.Equal(t, 1, r.ID)
		assert.Equal(t, 5, r.EntityID)
		assert.Equal(t, true, r.IsPrivate)
		assert.Equal(t, "Tyrion Lannister", r.Name)
		assert.Equal(t, "Lorem Ipsum", r.ToolTip)
		assert.Equal(t, "character", r.Type)
		assert.Equal(t, "https://example.com/campaign/1/characters/1", r.URL)
		assert.Equal(t, "https://example.com/image.png", r.Image)
		assert.Equal(t, "https://example.com/image_thumb.png", r.ImageThumb)
		assert.Equal(t, false, r.HasCustomImage)

		// Date & time assertions
		created := time.Date(2019, time.January, 30, 00, 01, 44, 0, time.UTC)
		updated := time.Date(2019, time.August, 29, 13, 48, 54, 0, time.UTC)
		assert.Equal(t, created, r.CreatedAt)
		assert.Equal(t, updated, r.UpdatedAt)
		assert.Equal(t, 1, r.CreatedBy)
		assert.Equal(t, 1, r.UpdatedBy)
	}
}

func TestResolveLinks(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	// Test successful resolution
	rawTestString := "[character:1] hails from [location:1]"
	resolvedTestString, err := client.Searches(1).ResolveLinks(ctx, rawTestString)

	if assert.NoError(t, err) {
		assert.Equal(t, "Jonathan Green hails from Mordor", resolvedTestString)
	}

	// Test unsuccessful resolution
	rawTestString = "[character:100] hails from [location:1]"
	resolvedTestString, err = client.Searches(1).ResolveLinks(ctx, rawTestString)

	if assert.NoError(t, err) {
		assert.Equal(t, "[character:100] hails from Mordor", resolvedTestString)
	}

}
