package notifier

import (
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/reminder"
	"github.com/lucasnevespereira/lembra/internal/repository"
	"github.com/robfig/cron"
	"log"
	"os/exec"
	"time"
)

type CronNotifier struct {
	repo *repository.ReminderRepository
	cron *cron.Cron
}

func NewCronNotifier() (*CronNotifier, error) {
	reminderRepository, err := repository.NewReminderRepository()
	if err != nil {
		return nil, err
	}
	return &CronNotifier{
		repo: reminderRepository,
		cron: cron.New(),
	}, nil
}

func (n *CronNotifier) Start() error {
	// cron to check every minute
	err := n.cron.AddFunc("@every 1m", func() {
		err := n.CheckReminders()
		if err != nil {
			fmt.Errorf("check reminders: %v", err)
		}
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

func (n *CronNotifier) CheckReminders() error {
	log.Println("Checking reminders...")
	reminders, err := n.repo.GetAll()
	if err != nil {
		return fmt.Errorf("repo.GetAll: %v", err)
	}

	now := time.Now().Format("2006-01-02T15:04")
	for _, r := range reminders {
		if r.Time == now {
			n.sendNotification(r)
		}
	}

	return nil
}

func (n *CronNotifier) sendNotification(reminder *reminder.Reminder) {
	// TODO: Replace with a cross-platform notification library
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s" sound name "%s"`, reminder.Message, reminder.Title, reminder.Sound))
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to send notification for reminder ID %s: %v \n", reminder.ID, err)
		return
	}
	log.Printf("Notification sent for reminder ID %s \n", reminder.ID)
	err = n.repo.DeleteByID(reminder.ID)
	if err != nil {
		log.Printf("Failed to delete reminder with id %s: %v \n", reminder.ID, err)
	}
}
