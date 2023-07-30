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
func (r *reactionRepository) Create(reaction model.Reaction) error {
	logger := logrus.WithFields(logrus.Fields{
		"reaction": utils.Dump(reaction),
	})

	err := r.db.Create(&reaction).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}

// FindMatch :nodoc:
func (r *reactionRepository) FindMatch(userBy, userTo string) (model.Reaction, error) {
	logger := logrus.WithFields(logrus.Fields{
		"user_by": userBy,
		"user_to": userTo,
	})

	var result model.Reaction

	qb := r.db.
		Where("user_by = ?", userBy).
		Where("user_to = ?", userTo).
		Where("type = ?", model.ReactionTypeLike)

	err := qb.Find(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return model.Reaction{}, err
	}

	return result, nil
}

// CreateMatched :nodoc:
func (r *reactionRepository) CreateMatched(reaction model.Reaction, matched model.Reaction) error {
	logger := logrus.WithFields(logrus.Fields{
		"reaction": utils.Dump(reaction),
		"matched":  utils.Dump(matched),
	})

	tx := r.db.Begin()

	err := tx.Create(&reaction).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}

	err = tx.Save(&matched).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}

	tx.Commit()
	return err
}
