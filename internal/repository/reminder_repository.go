package repository

import (
	"database/sql"
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/reminder"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type ReminderRepository struct {
	*sql.DB
}

func setupDatabase() (*sql.DB, error) {
	var err error

	// Open the SQLite database connection
	db, err := sql.Open("sqlite3", "reminders.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping reminders db: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS reminders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			message TEXT,
			sound TEXT,
			time TEXT
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create reminders table: %v", err)
	}

	return db, nil
}

func NewReminderRepository() (*ReminderRepository, error) {
	database, err := setupDatabase()
	if err != nil {
		return nil, err
	}
	return &ReminderRepository{
		database,
	}, nil
}

func (r *ReminderRepository) Save(reminder *reminder.Reminder) error {
	stmt, err := r.Prepare("INSERT INTO reminders (title, message, sound, time) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(reminder.Title, reminder.Message, reminder.Sound, reminder.Time)
	if err != nil {
		return fmt.Errorf("failed to execute SQL statement: %v", err)
	}

	return nil
}

func (r *ReminderRepository) GetByID(id string) (*reminder.Reminder, error) {
	row := r.QueryRow("SELECT title, message, sound, time FROM reminders WHERE id = ?", id)

	var title, message, sound string
	var timeStr string
	err := row.Scan(&title, &message, &sound, &timeStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("failed to scan reminder row: %v", err)
	}

	time, err := time.Parse("2006-01-02 15:04", timeStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse reminder time: %v", err)
	}

	return reminder.NewReminder(title, message, sound, time), nil
}
