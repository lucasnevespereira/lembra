package repository

type ReminderRepository interface {
	Create(reminder *ReminderDB) error
	GetAll() ([]*ReminderDB, error)
	GetByID(id string) (*ReminderDB, error)
	Update(reminder *ReminderDB) error
	UpdateNotified(id string, value bool) error
	DeleteByID(id string) error
	DeleteAll() error
	DeleteNotified() error
}

type ReminderDB struct {
	ID       string
	Title    string
	Message  string
	Time     string
	Notified bool
}

func (ReminderDB) TableName() string {
	return "reminders"
}
