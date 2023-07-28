package repository

import (
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type reactionRepository struct {
	db *gorm.DB
}

// NewReactionRepository :nodoc:
func NewReactionRepository(d *gorm.DB) model.ReactionRepository {
	return &reactionRepository{
		db: d,
	}
}

// Create :nodoc:
func (u *reactionRepository) Create(reaction *model.Reaction) error {
	logger := logrus.WithFields(logrus.Fields{
		"reaction": utils.Dump(reaction),
	})

	err := u.db.Create(reaction).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}
