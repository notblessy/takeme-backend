package repository

import (
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type subscriptionPlanRepository struct {
	db *gorm.DB
}

// NewSubscriptionPlanRepository :nodoc:
func NewSubscriptionPlanRepository(d *gorm.DB) model.SubscriptionPlanRepository {
	return &subscriptionPlanRepository{
		db: d,
	}
}

// BulkCreate :nodoc:
func (u *subscriptionPlanRepository) BulkCreate(subscriptionPlan []model.SubscriptionPlan) error {
	err := u.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&subscriptionPlan).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"subscriptionPlan": utils.Dump(subscriptionPlan),
		}).Error(err)
		return err
	}

	return err
}
