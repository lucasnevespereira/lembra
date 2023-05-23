package reminder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewReminder(t *testing.T) {
	title := "Meeting"
	message := "You have a meeting"
	time := "2023-05-25T10:00"
	reminder, err := NewReminder(title, message, time)
	assert.NoError(t, err)
	assert.NotNil(t, reminder)
	assert.NotEmpty(t, reminder.ID)
	assert.Equal(t, title, reminder.Title)
	assert.Equal(t, message, reminder.Message)
	assert.Equal(t, time, reminder.Time)
	assert.False(t, reminder.Notified)
}

func TestParseTime(t *testing.T) {
	testCases := []struct {
		name         string
		timeInput    string
		expectedTime string
	}{
		{
			name:         "Time input in HH:MM format",
			timeInput:    "10:00",
			expectedTime: getExpectedTime(time.Now(), "10:00"),
		},
		{
			name:         "Time input in full datetime format",
			timeInput:    "2023-05-25T10:00",
			expectedTime: "2023-05-25T10:00",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsedTime, err := ParseTime(tc.timeInput)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedTime, parsedTime)
		})
	}
}

func TestGenerateID(t *testing.T) {
	id, err := generateID()
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

// getExpectedTime string in "2006-01-02T15:04" format
func getExpectedTime(now time.Time, timeInput string) string {
	currentDate := now.Format("2006-01-02")
	return fmt.Sprintf("%sT%s", currentDate, timeInput)
}
