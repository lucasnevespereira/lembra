package notifier

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/lucasnevespereira/lembra/internal/pkg/reminder"
	"github.com/lucasnevespereira/lembra/internal/pkg/storage"
	"github.com/lucasnevespereira/lembra/internal/utils/logger"
	"github.com/lucasnevespereira/lembra/internal/utils/mapping"
	"github.com/robfig/cron"
	"time"
)

type CronNotifier struct {
	repo storage.ReminderStorage
	cron *cron.Cron
}

func NewCronNotifier(repository storage.ReminderStorage, cron *cron.Cron) *CronNotifier {
	return &CronNotifier{
		repo: repository,
		cron: cron,
	}
}

func (n *CronNotifier) Start() error {
	// cron to check every minute
	err := n.cron.AddFunc("@every 1m", func() {
		n.CheckReminders()
	})
	if err != nil {
		return fmt.Errorf("n.cron.AddFunc: %v", err)
	}
	logger.Log.Println("Cron job started")
	n.cron.Start()
	return nil
}

func (n *CronNotifier) Stop() {
	n.cron.Stop()
}

func (n *CronNotifier) CheckReminders() {
	logger.Log.Println("Checking reminders...")
	dbReminders, err := n.repo.GetAll()
	if err != nil {
		fmt.Errorf("repo.GetAll: %v", err)
	}

	reminders := mapping.ToRemindersDTO(dbReminders)
	now := time.Now().Format("2006-01-02T15:04")
	for _, r := range reminders {
		if r.Time == now {
			n.sendNotification(r)
		}
	}
}

func (n *CronNotifier) sendNotification(reminder *reminder.Reminder) {
	err := beeep.Alert(reminder.Title, reminder.Message, "assets/icon.png")
	if err != nil {
		fmt.Errorf("Failed to send notification for reminder ID %s: %v \n", reminder.ID, err)
		return
	}

	err = n.repo.UpdateNotified(reminder.ID, true)
	if err != nil {
		fmt.Errorf("Failed to update reminder with id %s: %v \n", reminder.ID, err)
	}

	fmt.Printf("Notification sent for reminder ID %s \n", reminder.ID)

	err = n.repo.DeleteByID(reminder.ID)
	if err != nil {
		fmt.Errorf("Failed to delete reminder with id %s: %v \n", reminder.ID, err)
	}
}
