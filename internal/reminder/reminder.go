package reminder

import (
	"fmt"
	"time"
)

type Reminder struct {
	ID      string
	Title   string
	Message string
	Sound   string
	Time    time.Time
}

func NewReminder(title, message, sound string, time time.Time) *Reminder {
	return &Reminder{
		Title:   title,
		Message: message,
		Sound:   sound,
		Time:    time,
	}
}

func ParseTime(timeStr string) (time.Time, error) {
	timeLayout := "15h04"
	parsedTime, err := time.Parse(timeLayout, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %v", err)
	}
	return parsedTime, nil
}
