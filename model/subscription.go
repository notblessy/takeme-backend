package model

import (
	"time"

	"gorm.io/datatypes"
)

// SubscriptionRepository :nodoc:
type SubscriptionRepository interface {
	Create(subscription *Subscription) error
	FindByID(userID string, subscription *Subscription) error
}

// SubscriptionUsecase :nodoc:
type SubscriptionUsecase interface {
	Create(subscription Subscription) (Subscription, error)
}

// SubscriptionPlanRepository :nodoc:
type SubscriptionPlanRepository interface {
	BulkCreate(subscriptionPlan []SubscriptionPlan) error
}

// SubscriptionPlan :nodoc:
type SubscriptionPlan struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Price     int            `json:"price"`
	Feature   datatypes.JSON `json:"feature"`
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at,omitempty"`
}

// Subscription :nodoc:
type Subscription struct {
	ID                 int              `json:"id"`
	UserID             string           `json:"user_id"`
	SubscriptionPlanID int              `json:"subscription_plan_id"`
	IsActive           bool             `json:"is_active"`
	CreatedAt          time.Time        `gorm:"<-:create" json:"created_at"`
	UpdatedAt          *time.Time       `json:"updated_at"`
	DeletedAt          *time.Time       `json:"deleted_at,omitempty"`
	SubscriptionPlan   SubscriptionPlan `gorm:"->" json:"subscription_plan"`
}

// Feature :nodoc:
type Feature struct {
	SwipeLimit    int  `json:"swipe_limit"`
	VerifiedBadge bool `json:"verified_badge"`
}

func NewDefaultSubscription(userID string) Subscription {
	return Subscription{
		UserID:             userID,
		SubscriptionPlanID: 1,
		IsActive:           true,
	}
}

func (s *Subscription) IsFree() bool {
	return s.SubscriptionPlan.Name == "FREE" || s.ID == 0
}
