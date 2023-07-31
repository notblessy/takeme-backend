package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/notblessy/takeme-backend/cacher"
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type reactionRepository struct {
	db    *gorm.DB
	cache cacher.Cacher
}

// NewReactionRepository :nodoc:
func NewReactionRepository(d *gorm.DB, c cacher.Cacher) model.ReactionRepository {
	return &reactionRepository{
		db:    d,
		cache: c,
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

	cacheKey := r.newReactionCacheKey(reaction.UserBy, reaction.UserTo)

	cacheItem := cacher.NewItemWithTTL(cacheKey, reaction, 24*time.Hour)
	err = r.cache.Store(cacheItem)
	if err != nil {
		logger.Error(err)
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

	// find match from cache
	cacheKey := r.newReactionCacheKey(userTo, userBy)
	err := r.findMatchFromCache(cacheKey, &result)
	if err != nil {
		logger.Error(err)
		return model.Reaction{}, err
	}

	if result.ID != "" {
		return result, nil
	}

	// find match from db if no cache
	qb := r.db.
		Where("user_by = ?", userBy).
		Where("user_to = ?", userTo).
		Where("type = ?", model.ReactionTypeLike)

	err = qb.Find(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return model.Reaction{}, err
	}

	cacheItem := cacher.NewItemWithTTL(cacheKey, result, 24*time.Hour)
	err = r.cache.Store(cacheItem)
	if err != nil {
		logger.Error(err)
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

// FindAllSwiped :nodoc:
func (r *reactionRepository) FindAllSwiped(userBy string) ([]string, error) {
	logger := logrus.WithFields(logrus.Fields{
		"user_by": userBy,
	})

	var res []string

	err := r.db.Table("reactions").Select("user_to").Where("user_by = ?", userBy).Scan(&res).Error
	if err != nil {
		logger.Error(err)
		return []string{}, err
	}

	return res, nil
}

func (r *reactionRepository) findMatchFromCache(cacheKey string, reaction *model.Reaction) error {
	res, err := r.cache.Get(cacheKey)
	if err != nil {
		return err
	}

	switch res {
	case nil:
		return nil
	default:
		err := json.Unmarshal(res.([]byte), reaction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *reactionRepository) newReactionCacheKey(userBy, userTo string) string {
	return fmt.Sprintf("match:%s:%s", userBy, userTo)
}
