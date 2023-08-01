package repository

import (
	"time"

	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type subscriptionRepository struct {
	db *gorm.DB
}

// NewSubscriptionRepository :nodoc:
func NewSubscriptionRepository(d *gorm.DB) model.SubscriptionRepository {
	return &subscriptionRepository{
		db: d,
	}
}

// Create :nodoc:
func (u *subscriptionRepository) Create(subscription *model.Subscription) error {
	logger := logrus.WithFields(logrus.Fields{
		"subscription": utils.Dump(subscription),
	})

	tx := u.db.Begin()

	err := tx.Table("subscriptions").Where("user_id = ?", subscription.UserID).Updates(map[string]interface{}{
		"is_active":  false,
		"deleted_at": time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}

	err = tx.Create(subscription).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}

	tx.Commit()
	return err
}

// FindByID :nodoc:
func (u *subscriptionRepository) FindByID(userID string, subscription *model.Subscription) error {
	logger := logrus.WithFields(logrus.Fields{
		"user_id": userID,
	})

	err := u.db.Joins("SubscriptionPlan").
		Where("subscriptions.user_id = ?", userID).
		Where("subscriptions.is_active = ?", true).
		First(subscription).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
