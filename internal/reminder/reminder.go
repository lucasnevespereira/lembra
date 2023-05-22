package reminder

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Reminder struct {
	ID        string
	Title     string
	Message   string
	Sound     string
	Time      time.Time
	Scheduled bool
	Notified  bool
}

func NewReminder(title, message, sound string, time time.Time) (*Reminder, error) {
	id, err := generateID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate ID: %v", err)
	}
	return &Reminder{
		ID:        id,
		Title:     title,
		Message:   message,
		Sound:     sound,
		Time:      time,
		Scheduled: false,
		Notified:  false,
	}, nil
}

func ParseTime(timeStr string) (time.Time, error) {
	timeLayout := "2006-01-02 15h04"
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	parsedTime, err := time.ParseInLocation(timeLayout, fmt.Sprintf("%s %s", currentDate, timeStr), currentTime.Location())
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %v", err)
	}
	return parsedTime, nil
}

func generateID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %v", err)
	}
	return id.String(), nil
}
