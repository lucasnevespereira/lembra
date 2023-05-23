package repository

import (
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/pkg/reminder"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ReminderRepository struct {
	db *gorm.DB
}

func NewReminderRepository() (*ReminderRepository, error) {
	db, err := gorm.Open(sqlite.Open("reminders.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	err = db.AutoMigrate(&reminder.Reminder{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate reminders table: %v", err)
	}

	return &ReminderRepository{
		db: db,
	}, nil
}

func (r *ReminderRepository) Create(reminder *reminder.Reminder) error {
	result := r.db.Create(reminder)
	if result.Error != nil {
		return fmt.Errorf("failed to create reminder: %v", result.Error)
	}
	return nil
}

func (r *ReminderRepository) GetAll() ([]*reminder.Reminder, error) {
	var reminders []*reminder.Reminder
	result := r.db.Find(&reminders)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve reminders: %v", result.Error)
	}
	return reminders, nil
}

func (r *ReminderRepository) GetByID(id string) (*reminder.Reminder, error) {
	var reminder *reminder.Reminder
	result := r.db.First(&reminder, "id = ?", id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve reminder with id %s: %v", id, result.Error)
	}

	return reminder, nil
}

func (r *ReminderRepository) Update(reminder *reminder.Reminder) error {
	result := r.db.Save(reminder)
	if result.Error != nil {
		return fmt.Errorf("failed to update reminder: %v", result.Error)
	}
	return nil
}

func (r *ReminderRepository) UpdateNotified(reminder *reminder.Reminder, value bool) error {
	result := r.db.Model(&reminder).Update("notified", value)
	if result.Error != nil {
		return fmt.Errorf("failed to update reminder: %v", result.Error)
	}
	return nil
}

func (r *ReminderRepository) DeleteByID(id string) error {
	result := r.db.Delete(&reminder.Reminder{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete reminder with id %s: %v", id, result.Error)
	}
	return nil
}

func (r *ReminderRepository) DeleteAll() error {
	result := r.db.Where("id not null").Delete(&reminder.Reminder{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete all reminders: %v", result.Error)
	}
	return nil
}

func (r *ReminderRepository) DeleteNotified() error {
	result := r.db.Delete(&reminder.Reminder{}, "notified = ?", true)
	if result.Error != nil {
		return fmt.Errorf("failed to delete notified reminders: %v", result.Error)
	}
	return nil
}
