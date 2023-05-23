package repository

import (
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/utils/logger"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewReminderRepository(db *gorm.DB) ReminderRepository {
	err := db.AutoMigrate(&ReminderDB{})
	if err != nil {
		logger.Log.Printf("failed to migrate reminders table: %v", err)
	}
	return &repository{
		db: db,
	}
}

func (r *repository) Create(reminder *ReminderDB) error {
	result := r.db.Create(reminder)
	if result.Error != nil {
		return fmt.Errorf("failed to create reminder: %v", result.Error)
	}
	return nil
}

func (r *repository) GetAll() ([]*ReminderDB, error) {
	var reminders []*ReminderDB
	result := r.db.Find(&reminders)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve reminders: %v", result.Error)
	}
	return reminders, nil
}

func (r *repository) GetByID(id string) (*ReminderDB, error) {
	var reminder *ReminderDB
	result := r.db.First(&reminder, "id = ?", id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve reminder with id %s: %v", id, result.Error)
	}

	return reminder, nil
}

func (r *repository) Update(reminder *ReminderDB) error {
	result := r.db.Save(reminder)
	if result.Error != nil {
		return fmt.Errorf("failed to update reminder: %v", result.Error)
	}
	return nil
}

func (r *repository) UpdateNotified(id string, value bool) error {
	result := r.db.Model(&ReminderDB{}).Where("id = ?", id).Update("notified", value)
	if result.Error != nil {
		return fmt.Errorf("failed to update reminder: %v", result.Error)
	}
	return nil
}

func (r *repository) DeleteByID(id string) error {
	result := r.db.Delete(&ReminderDB{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete reminder with id %s: %v", id, result.Error)
	}
	return nil
}

func (r *repository) DeleteAll() error {
	result := r.db.Where("id not null").Delete(&ReminderDB{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete all reminders: %v", result.Error)
	}
	return nil
}

func (r *repository) DeleteNotified() error {
	result := r.db.Delete(&ReminderDB{}, "notified = ?", true)
	if result.Error != nil {
		return fmt.Errorf("failed to delete notified reminders: %v", result.Error)
	}
	return nil
}
