package notifier

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/lucasnevespereira/lembra/internal/pkg/reminder"
	"github.com/lucasnevespereira/lembra/internal/pkg/repository"
	"github.com/robfig/cron"
	"log"
	"time"
)

type CronNotifier struct {
	repo *repository.ReminderRepository
	cron *cron.Cron
}

func NewCronNotifier(repository *repository.ReminderRepository, cron *cron.Cron) *CronNotifier {
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
	log.Println("Cron job started")
	n.cron.Start()
	return nil
}

func (n *CronNotifier) Stop() {
	n.cron.Stop()
}

func (n *CronNotifier) CheckReminders() {
	log.Println("Checking reminders...")
	reminders, err := n.repo.GetAll()
	if err != nil {
		log.Printf("repo.GetAll: %v", err)
	}

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
		log.Printf("Failed to send notification for reminder ID %s: %v \n", reminder.ID, err)
		return
	}

	err = n.repo.UpdateNotified(reminder, true)
	if err != nil {
		log.Printf("Failed to update reminder with id %s: %v \n", reminder.ID, err)
	}

	log.Printf("Notification sent for reminder ID %s \n", reminder.ID)

	err = n.repo.DeleteByID(reminder.ID)
	if err != nil {
		log.Printf("Failed to delete reminder with id %s: %v \n", reminder.ID, err)
	}
}
