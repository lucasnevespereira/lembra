package mapping

import (
	"github.com/lucasnevespereira/lembra/internal/pkg/reminder"
	"github.com/lucasnevespereira/lembra/internal/pkg/repository"
)

func ToReminderDTO(reminderDB *repository.ReminderDB) *reminder.Reminder {
	return &reminder.Reminder{
		ID:       reminderDB.ID,
		Title:    reminderDB.Title,
		Message:  reminderDB.Message,
		Time:     reminderDB.Time,
		Notified: reminderDB.Notified,
	}
}

func ToRemindersDTO(reminderDBs []*repository.ReminderDB) []*reminder.Reminder {
	reminders := make([]*reminder.Reminder, len(reminderDBs))
	for i, reminderDB := range reminderDBs {
		reminders[i] = ToReminderDTO(reminderDB)
	}
	return reminders
}
