package storage

type ReminderStorage interface {
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
	ID       string `json:ID`
	Title    string `json:"Title"`
	Message  string `json:"Message"`
	Time     string `json:"Time"`
	Notified bool   `json:"Notified"`
}
