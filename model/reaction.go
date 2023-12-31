package model

import (
	"time"
)

const (
	ReactionTypeLike string = "LIKE"
	ReactionTypePass string = "PASS"
)

// ReactionRepository :nodoc:
type ReactionRepository interface {
	Create(reaction Reaction) error
	CreateMatched(reaction Reaction, matched Reaction) error
	FindMatch(userBy, userTo string) (Reaction, error)
	FindAllSwiped(userBy string) ([]string, error)
	FindTotalSwipeToday(userBy string) (int64, error)
}

// ReactionUsecase :nodoc:
type ReactionUsecase interface {
	Create(reaction ReactionRequest) (Reaction, error)
}

// Reaction :nodoc:
type Reaction struct {
	ID        string     `json:"id"`
	UserBy    string     `json:"user_by"`
	UserTo    string     `json:"user_to"`
	Type      string     `json:"type"`
	MatchedAt *time.Time `json:"matched_at"`
	CreatedAt time.Time  `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// ReactionRequest :nodoc:
type ReactionRequest struct {
	UserBy string `json:"user_by"`
	UserTo string `json:"user_to" validate:"required"`
	Type   string `json:"type" validate:"required"`
}

// NewReaction :nodoc:
func NewReaction(request ReactionRequest) Reaction {
	return Reaction{
		UserBy: request.UserBy,
		UserTo: request.UserTo,
		Type:   request.Type,
	}
}

// IsMatch :nodoc:
func (r *Reaction) IsMatch() bool {
	return r.ID != ""
}
