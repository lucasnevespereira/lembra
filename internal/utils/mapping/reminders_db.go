package mapping

import (
	"github.com/lucasnevespereira/lembra/internal/pkg/reminder"
	"github.com/lucasnevespereira/lembra/internal/pkg/repository"
)

func ToReminderDB(reminder *reminder.Reminder) *repository.ReminderDB {
	return &repository.ReminderDB{
		ID:       reminder.ID,
		Title:    reminder.Title,
		Message:  reminder.Message,
		Time:     reminder.Time,
		Notified: reminder.Notified,
	}
}
