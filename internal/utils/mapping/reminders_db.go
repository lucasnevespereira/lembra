package mapping

import (
	"github.com/lucasnevespereira/lembra/internal/pkg/reminder"
	"github.com/lucasnevespereira/lembra/internal/pkg/storage"
)

func ToReminderDB(reminder *reminder.Reminder) *storage.ReminderDB {
	return &storage.ReminderDB{
		ID:       reminder.ID,
		Title:    reminder.Title,
		Message:  reminder.Message,
		Time:     reminder.Time,
		Notified: reminder.Notified,
	}
}
