package zoom

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// CreateMeetingOptions are the options to create a meeting with
type CreateMeetingOptions struct {
	HostID         string          `json:"-"`
	Topic          string          `json:"topic,omitempty"`
	Type           MeetingType     `json:"type,omitempty"`
	StartTime      *time.Time      `json:"start_time,omitempty"`
	Duration       int             `json:"duration,omitempty"`
	Timezone       string          `json:"timezone,omitempty"`
	Password       string          `json:"password,omitempty"` // Max 10 characters. [a-z A-Z 0-9 @ - _ *]
	Agenda         string          `json:"agenda,omitempty"`
	TrackingFields []TrackingField `json:"tracking_fields,omitempty"`
	Settings       MeetingSettings `json:"settings,omitempty"`
}

// Meeting represents a meeting created/returned by GetMeeting
type Meeting struct {
	UUID              string        `json:"uuid,omitempty"`
	ID                int           `json:"id,omitempty"`
	HostID            string        `json:"host_id,omitempty"`
	Topic             string        `json:"topic"`
	Type              MeetingType   `json:"type"`
	Status            MeetingStatus `json:"status"`
	StartTime         *time.Time    `json:"start_time"`
	Duration          int           `json:"duration"`
	Timezone          string        `json:"timezone"`
	CreatedAt         *time.Time    `json:"created_at"`
	Agenda            string        `json:"agenda"`
	StartURL          string        `json:"start_url"`
	JoinURL           string        `json:"join_url"`
	Password          string        `json:"password"`
	H323Password      string        `json:"h323_password"`
	EncryptedPassword string        `json:"encrypted_password"`
	// PMI is Personal Meeting ID. Only used for scheduled meetings and recurring meetings with
	// no fixed time
	PMI            int64           `json:"pmi"`
	TrackingFields []TrackingField `json:"tracking_fields"`
	Occurrences    []Occurrence    `json:"occurrences"`
	Settings       MeetingSettings `json:"settings"`
	Recurrence     Recurrence      `json:"recurrence"`
}

// Recurrence of the meeting
type Recurrence struct {
	Type           RecurrenceType `json:"type"`
	RepeatInterval int            `json:"repeat_interval"`
	WeeklyDays     string         `json:"weekly_days"`
	MonthlyDay     int            `json:"monthly_day"`
	MonthlyWeek    MonthlyWeek    `json:"monthly_week"`
	MonthlyWeekDay WeekDay        `json:"monthly_week_day"`
	// EndTimes how many times the meeting will recur before it is canceled (cannot be used
	// with "end_time_date"
	EndTimes int `json:"end_times"`
	// EndDateTime should be in UTC. Cannot be used with "end_times"
	EndDateTime *time.Time `json:"end_date_time"`
}

type MeetingsData struct {
	PageSize      int       `json:"page_size,omitempty"`
	TotalRecords  int       `json:"total_records,omitempty"`
	NextPageToken string    `json:"next_page_token,omitempty"`
	Meetings      []Meeting `json:"meetings,omitempty"`
}

func (c *Client) Meetings() (*MeetingsData, error) {
	body, err := c.createRequest("/users/me/meetings", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	meetingData := MeetingsData{}
	json.Unmarshal(*body, &meetingData)
	return &meetingData, nil
}

func (c *Client) CreateMeeting(createOption CreateMeetingOptions) (*Meeting, error) {
	jd, err := json.Marshal(createOption)
	if err != nil {
		return nil, err
	}
	data := bytes.NewBuffer(jd)
	body, err := c.createRequest("/users/me/meetings", http.MethodPost, data)
	if err != nil {
		return nil, err
	}
	meeting := &Meeting{}
	json.Unmarshal(*body, meeting)
	return meeting, nil
}
