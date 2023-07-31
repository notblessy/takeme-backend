package model

import (
	"time"
)

var SupportedGender = map[string]int{
	"MALE":   1,
	"FEMALE": 2,
	"BOTH":   3,
}

// UserRepository :nodoc:
type UserRepository interface {
	Create(User User) error
	FindByEmail(email string) (User, error)
	FindByID(id string, user *User) error
	FindAll(request map[string]string, userIDs []string) ([]User, int64, error)
}

// UserUsecase :nodoc:
type UserUsecase interface {
	Register(User RegisterUser) (User, error)
	Login(user User) (string, error)

	FindAll(request map[string]string, userID string) ([]User, int64, error)
}

// User :nodoc:
type User struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email" validate:"required"`
	Password    string     `json:"password,omitempty"`
	Description string     `json:"description"`
	Gender      int        `json:"gender"`
	Preference  int        `json:"preference"`
	Age         int        `json:"age"`
	IsPremium   bool       `json:"is_premium"`
	CreatedAt   time.Time  `gorm:"<-:create" json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type RegisterUser struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email" validate:"required"`
	Password    string   `json:"password,omitempty"`
	Description string   `json:"description"`
	Gender      string   `json:"gender"`
	Preference  string   `json:"preference"`
	Age         int      `json:"age"`
	Photos      []string `json:"photos,omitempty"`
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

func (u *User) NewUserFromRequest(req RegisterUser) {
	u.Name = req.Name
	u.Description = req.Description
	u.Email = req.Email
	u.Password = req.Password
	u.Gender = SupportedGender[req.Gender]
	u.Preference = SupportedGender[req.Preference]
	u.Age = req.Age
}

func (u *User) IsPasswordCorrect(req User) bool {
	if u.Email == req.Email && u.Password == req.Password {
		return true
	}

	return false
}
