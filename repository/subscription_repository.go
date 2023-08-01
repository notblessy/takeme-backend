package repository

import (
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
	err := u.db.Create(subscription).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"subscription": utils.Dump(subscription),
		}).Error(err)
		return err
	}

	return err
}
