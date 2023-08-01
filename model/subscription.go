package model

import (
	"time"

	"gorm.io/datatypes"
)

// SubscriptionRepository :nodoc:
type SubscriptionRepository interface {
	Create(subscription *Subscription) error
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
	ID                 int        `json:"id"`
	UserID             string     `json:"user_id"`
	SubscriptionPlanID string     `json:"subscription_plan_id"`
	IsActive           bool       `json:"is_active"`
	ExpiredAt          *time.Time `json:"expired_at"`
	CreatedAt          time.Time  `gorm:"<-:create" json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}

// Feature :nodoc:
type Feature struct {
	SwipeLimit    int  `json:"swipe_limit"`
	VerifiedBadge bool `json:"verified_badge"`
}
