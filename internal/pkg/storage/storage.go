package storage

import (
	"encoding/json"
	"os"
)

type storage struct {
	filePath string
}

func NewReminderStorage(filePath string) ReminderStorage {
	return &storage{
		filePath: filePath,
	}
}

func (s *storage) Create(reminder *ReminderDB) error {
	reminders, err := s.GetAll()
	if err != nil {
		return err
	}

	reminders = append(reminders, reminder)

	data, err := json.MarshalIndent(reminders, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

func (s *storage) GetAll() ([]*ReminderDB, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var reminders []*ReminderDB
	err = json.Unmarshal(data, &reminders)
	if err != nil {
		return nil, err
	}

	return reminders, nil
}

func (s *storage) GetByID(id string) (*ReminderDB, error) {
	reminders, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	for _, reminder := range reminders {
		if reminder.ID == id {
			return reminder, nil
		}
	}

	return nil, nil
}

func (s *storage) Update(reminder *ReminderDB) error {
	reminders, err := s.GetAll()
	if err != nil {
		return err
	}

	for i, existingReminder := range reminders {
		if existingReminder.ID == reminder.ID {
			reminders[i] = reminder
			data, err := json.MarshalIndent(reminders, "", "    ")
			if err != nil {
				return err
			}

			return os.WriteFile(s.filePath, data, 0644)
		}
	}

	return nil
}

func (s *storage) UpdateNotified(id string, value bool) error {
	reminders, err := s.GetAll()
	if err != nil {
		return err
	}

	for _, reminder := range reminders {
		if reminder.ID == id {
			reminder.Notified = value
			data, err := json.MarshalIndent(reminders, "", "    ")
			if err != nil {
				return err
			}

			return os.WriteFile(s.filePath, data, 0644)
		}
	}

	return nil
}

func (s *storage) DeleteByID(id string) error {
	reminders, err := s.GetAll()
	if err != nil {
		return err
	}

	var updatedReminders []*ReminderDB
	for _, reminder := range reminders {
		if reminder.ID != id {
			updatedReminders = append(updatedReminders, reminder)
		}
	}

	data, err := json.MarshalIndent(updatedReminders, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

func (s *storage) DeleteAll() error {
	// Simply create an empty file to delete all reminders
	return os.WriteFile(s.filePath, []byte("[]"), 0644)
}

func (s *storage) DeleteNotified() error {
	reminders, err := s.GetAll()
	if err != nil {
		return err
	}

	var updatedReminders []*ReminderDB
	for _, reminder := range reminders {
		if !reminder.Notified {
			updatedReminders = append(updatedReminders, reminder)
		}
	}

	data, err := json.MarshalIndent(updatedReminders, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}
