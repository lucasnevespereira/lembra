package reminder

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Reminder struct {
	ID       string
	Title    string
	Message  string
	Time     string
	Notified bool
}

func NewReminder(title, message, time string) (*Reminder, error) {
	id, err := generateID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate ID: %v", err)
	}
	return &Reminder{
		ID:       id,
		Title:    title,
		Message:  message,
		Time:     time,
		Notified: false,
	}, nil
}

func ParseTime(timeInput string) (string, error) {
	timeLayout := "2006-01-02T15:04"

	// Check if timeInput has the format "15:04"
	if len(timeInput) == 5 && timeInput[2] == ':' {
		timeInput = fmt.Sprintf("%sT%s", time.Now().Format("2006-01-02"), timeInput)
	}

	parsedTime, err := time.Parse(timeLayout, timeInput)
	if err != nil {
		return "", fmt.Errorf("failed to parse time: %v", err)
	}

	formatted := parsedTime.Format(timeLayout)
	return formatted, nil
}

func generateID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %v", err)
	}
	return id.String(), nil
}
