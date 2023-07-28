package model

import (
	"time"
)

// UserRepository :nodoc:
type UserRepository interface {
	Create(User User) error
	FindByEmail(email string) (User, error)
	FindByID(id string, user *User) error
	FindAllUsersByRole(organizationID, role string, user *[]User) error
}

// UserUsecase :nodoc:
type UserUsecase interface {
	Register(User User) (User, error)
	Login(user User) (string, error)
}

// User :nodoc:
type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email" validate:"required"`
	Password  string     `json:"password,omitempty"`
	Address   string     `json:"address"`
	Photo     string     `json:"photo"`
	CreatedAt time.Time  `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// AuthRequest :nodoc:
type AuthRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Auth :nodoc:
type Auth struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (u *User) IsPasswordCorrect(req User) bool {
	if u.Email == req.Email && u.Password != req.Password {
		return false
	}

	return true
}
