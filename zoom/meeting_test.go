package zoom

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateMeeting(t *testing.T) {
	t.Run("It sould create meeting", func(t *testing.T) {
		c := NewClient(apiKey, apiSecret)
		tm := time.Now().Add(time.Hour * 4)
		options := CreateMeetingOptions{
			Topic:     "Test meeting 1",
			Agenda:    "Test meeting create",
			StartTime: &tm,
			Duration:  30,
			Type:      MeetingTypeScheduled,
			Password:  "s2r23j",
		}
		got, err := c.CreateMeeting(options)
		assert.Nil(t, err, "Error should be nil")
		assert.NotNil(t, got, "Meeting can not be nil")
		assert.Equal(t, got.Topic, options.Topic, "Topic should be equal")
		assert.Equal(t, got.Password, options.Password, "Topic should be equal")
	})
	t.Run("It sould fail to create meeting due to invalid api key", func(t *testing.T) {
		c := &Client{
			apiKey:    "a1d2m4p5i_k3m5e2y0y",
			endpoint:  API_ENDPOINT,
			secretKey: apiSecret,
		}
		tm := time.Now().Add(time.Hour * 4)
		options := CreateMeetingOptions{
			Topic:     "Test meeting 1",
			Agenda:    "Test meeting create",
			StartTime: &tm,
			Duration:  30,
			Type:      MeetingTypeScheduled,
			Password:  "s2r23j",
		}
		meet, err := c.CreateMeeting(options)
		assert.NotNil(t, err, "Error should not be nil")
		assert.Nil(t, meet, "Meeting data should be nil")
	})
}
