package kanka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalendars(t *testing.T) {

	client := NewClient(DefaultConfig())
	c := client.Calendars(1)

	assert.Equal(t, client, c.client)
	assert.Equal(t, "/campaigns/1/calendars", c.urlPrefix)
}

func TestGetCalendars(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	calendars, err := client.Calendars(1).GetCalendars(ctx)

	if assert.NoError(t, err) {
		assert.Len(t, *calendars, 1)

		// Main calendars assertions
		c := (*calendars)[0]
		assert.Equal(t, 1, c.ID)
		assert.Equal(t, 78, c.EntityID)
		assert.Equal(t, true, c.IsPrivate)
		assert.Equal(t, "Georgian Calendar", c.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", c.Entry)
		assert.Equal(t, "311-2-3", c.Date)
		assert.Equal(t, "Primary", c.Type)
		assert.Equal(t, "https://example.com/image.png", c.Image)
		assert.Equal(t, "https://example.com/image_full.png", c.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", c.ImageThumb)
		assert.Equal(t, false, c.HasCustomImage)

		// Divisions
		assert.Len(t, c.Weekdays, 7)
		assert.Equal(t, "Mol", c.Weekdays[1])

		assert.Len(t, c.Months, 2)
		assert.Equal(t, "February", c.Months[1].Name)
		assert.Equal(t, 5, c.Months[1].Length)
		assert.Equal(t, "intercalary", c.Months[1].Type)

		assert.Len(t, c.Years, 2)
		assert.Equal(t, "Year of Water and Bone", c.Years["300"])

		assert.Len(t, c.Moons, 2)
		assert.Equal(t, "Olarune", c.Moons[1].Name)
		assert.Equal(t, "17", c.Moons[1].Fullmoon)
		assert.Equal(t, 1, c.Moons[1].Offset)
		assert.Equal(t, "brown", c.Moons[1].Colour)

		assert.Len(t, c.Seasons, 2)
		assert.Equal(t, "Summer", c.Seasons[1].Name)
		assert.Equal(t, 4, c.Seasons[1].Month)
		assert.Equal(t, 1, c.Seasons[1].Day)

		assert.Equal(t, 1, c.StartOffset)
		assert.Equal(t, "BC", c.Suffix)
		assert.Equal(t, true, c.HasLeapYear)
		assert.Equal(t, 4, c.LeapYearAmount)
		assert.Equal(t, 2, c.LeapYearMonth)
		assert.Equal(t, 3, c.LeapYearOffset)
		assert.Equal(t, 233, c.LeapYearStart)

		// Date & time assertions
		created := time.Date(2019, time.January, 28, 6, 29, 29, 0, time.UTC)
		updated := time.Date(2020, time.January, 30, 17, 30, 52, 0, time.UTC)
		assert.Equal(t, created, c.CreatedAt)
		assert.Equal(t, updated, c.UpdatedAt)
		assert.Equal(t, 1, c.CreatedBy)
		assert.Equal(t, 1, c.UpdatedBy)
	}
}

func TestGetCalendar(t *testing.T) {

	testServer, config := mockTestServer()
	defer testServer.Close()
	client := NewClient(config)
	ctx := context.Background()

	calendar, err := client.Calendars(1).GetCalendar(ctx, 1)

	if assert.NoError(t, err) {

		// Main calendar assertions
		c := *calendar
		assert.Equal(t, 1, c.ID)
		assert.Equal(t, 78, c.EntityID)
		assert.Equal(t, true, c.IsPrivate)
		assert.Equal(t, "Georgian Calendar", c.Name)
		assert.Equal(t, "\n<p>Lorem Ipsum.</p>\n", c.Entry)
		assert.Equal(t, "311-2-3", c.Date)
		assert.Equal(t, "Primary", c.Type)
		assert.Equal(t, "https://example.com/image.png", c.Image)
		assert.Equal(t, "https://example.com/image_full.png", c.ImageFull)
		assert.Equal(t, "https://example.com/image_thumb.png", c.ImageThumb)
		assert.Equal(t, false, c.HasCustomImage)

		// Divisions
		assert.Len(t, c.Weekdays, 7)
		assert.Equal(t, "Mol", c.Weekdays[1])

		assert.Len(t, c.Months, 2)
		assert.Equal(t, "February", c.Months[1].Name)
		assert.Equal(t, 5, c.Months[1].Length)
		assert.Equal(t, "intercalary", c.Months[1].Type)

		assert.Len(t, c.Years, 2)
		assert.Equal(t, "Year of Water and Bone", c.Years["300"])

		assert.Len(t, c.Moons, 2)
		assert.Equal(t, "Olarune", c.Moons[1].Name)
		assert.Equal(t, "17", c.Moons[1].Fullmoon)
		assert.Equal(t, 1, c.Moons[1].Offset)
		assert.Equal(t, "brown", c.Moons[1].Colour)

		assert.Len(t, c.Seasons, 2)
		assert.Equal(t, "Summer", c.Seasons[1].Name)
		assert.Equal(t, 4, c.Seasons[1].Month)
		assert.Equal(t, 1, c.Seasons[1].Day)

		assert.Equal(t, 1, c.StartOffset)
		assert.Equal(t, "BC", c.Suffix)
		assert.Equal(t, true, c.HasLeapYear)
		assert.Equal(t, 4, c.LeapYearAmount)
		assert.Equal(t, 2, c.LeapYearMonth)
		assert.Equal(t, 3, c.LeapYearOffset)
		assert.Equal(t, 233, c.LeapYearStart)

		// Date & time assertions
		created := time.Date(2019, time.January, 28, 6, 29, 29, 0, time.UTC)
		updated := time.Date(2020, time.January, 30, 17, 30, 52, 0, time.UTC)
		assert.Equal(t, created, c.CreatedAt)
		assert.Equal(t, updated, c.UpdatedAt)
		assert.Equal(t, 1, c.CreatedBy)
		assert.Equal(t, 1, c.UpdatedBy)
	}
}
