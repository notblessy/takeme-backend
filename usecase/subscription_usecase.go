package usecase

import (
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
)

type subscriptionUsecase struct {
	subscriptionRepo model.SubscriptionRepository
}

// NewSubscriptionUsecase :nodoc:
func NewSubscriptionUsecase(r model.SubscriptionRepository) model.SubscriptionUsecase {
	return &subscriptionUsecase{
		subscriptionRepo: r,
	}
}

// Create :nodoc:
func (u *subscriptionUsecase) Create(subscription model.Subscription) (model.Subscription, error) {
	logger := logrus.WithFields(logrus.Fields{
		"subscription": utils.Dump(subscription),
	})

	err := u.subscriptionRepo.Create(&subscription)
	if err != nil {
		logger.Error(err.Error())
		return model.Subscription{}, err
	}

	return subscription, nil
}
