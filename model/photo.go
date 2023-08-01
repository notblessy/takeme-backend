package model

import "time"

// Photo :nodoc:
type Photo struct {
	ID        int        `json:"id"`
	UserID    string     `json:"user_id"`
	URL       string     `json:"url"`
	CreatedAt time.Time  `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
