package model

import (
	"time"
)

// SupportedGender :nodoc:
var SupportedGender = map[string]int{
	"MALE":   1,
	"FEMALE": 2,
	"BOTH":   3,
}

// GenderMapper :nodoc:
var GenderMapper = map[int]string{
	1: "MALE",
	2: "FEMALE",
	3: "BOTH",
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

	FindAll(request map[string]string, userID string) ([]UserResponse, int64, error)
}

// User :nodoc:
type User struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email" validate:"required"`
	Password    string     `gorm:"->:false;<-:create" json:"password,omitempty"`
	Description string     `json:"description"`
	Gender      int        `json:"gender"`
	Preference  int        `json:"preference"`
	Age         int        `json:"age"`
	Photos      []Photo    `json:"photos"`
	CreatedAt   time.Time  `gorm:"<-:create" json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// RegisterUser :nodoc:
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

// UserResponse :nodoc:
type UserResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email" validate:"required"`
	Password    string     `gorm:"->:false;<-:create" json:"password,omitempty"`
	Description string     `json:"description"`
	Gender      string     `json:"gender"`
	Preference  string     `json:"preference"`
	Age         int        `json:"age"`
	IsPremium   bool       `json:"is_premium"`
	Photos      []Photo    `json:"photos"`
	CreatedAt   time.Time  `gorm:"<-:create" json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// NewUserFromRequest :nodoc:
func (u *User) NewUserFromRequest(req RegisterUser) {
	u.Name = req.Name
	u.Description = req.Description
	u.Email = req.Email
	u.Password = req.Password
	u.Gender = SupportedGender[req.Gender]
	u.Preference = SupportedGender[req.Preference]
	u.Age = req.Age
}

// NewUserPhotos :nodoc:
func (u *User) NewUserPhotos(photos []string) {
	for _, url := range photos {
		p := Photo{
			UserID: u.ID,
			URL:    url,
		}

		u.Photos = append(u.Photos, p)
	}
}

// IsPasswordCorrect :nodoc:
func (u *User) IsPasswordCorrect(req User) bool {
	if u.Email == req.Email && u.Password == req.Password {
		return true
	}

	return false
}

// ToResponse :nodoc:
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Description: u.Description,
		Gender:      GenderMapper[u.Gender],
		Preference:  GenderMapper[u.Preference],
		Age:         u.Age,
		Photos:      u.Photos,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		DeletedAt:   u.DeletedAt,
	}
}
