package repository

import (
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

// NewNotificationRepository :nodoc:
func NewNotificationRepository(d *gorm.DB) model.NotificationRepository {
	return &notificationRepository{
		db: d,
	}
}

// Create :nodoc:
func (u *notificationRepository) Create(notification model.Notification) error {
	err := u.db.Create(notification).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"notification": utils.Dump(notification),
		}).Error(err)
		return err
	}

	return err
}
