package model

import "time"

// NotificationRepository :nodoc:
type NotificationRepository interface {
	CreateInBatch(notif []Notification) error
}

// Notification :nodoc:
type Notification struct {
	ID        string     `json:"id"`
	UserBy    string     `json:"user_id"`
	UserTo    string     `json:"user_to"`
	Message   string     `json:"message"`
	IsRead    bool       `json:"is_read"`
	CreatedAt time.Time  `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewMatchNotification(reaction Reaction, message string) []Notification {
	notif1 := Notification{
		UserBy:  reaction.UserBy,
		UserTo:  reaction.UserTo,
		Message: message,
	}

	notif2 := Notification{
		UserBy:  reaction.UserTo,
		UserTo:  reaction.UserBy,
		Message: message,
	}

	return []Notification{notif1, notif2}
}
