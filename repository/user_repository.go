package repository

import (
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository :nodoc:
func NewUserRepository(d *gorm.DB) model.UserRepository {
	return &userRepository{
		db: d,
	}
}

// Create :nodoc:
func (u *userRepository) Create(user model.User) error {
	logger := logrus.WithFields(logrus.Fields{
		"user": utils.Dump(user),
	})

	tx := u.db.Begin()

	err := tx.Table("users").Create(&user).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}

	defaultSub := model.NewDefaultSubscription(user.ID)

	err = tx.Create(&defaultSub).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}

	tx.Commit()

	return err
}

// FindByEmail :nodoc:
func (u *userRepository) FindByEmail(email string) (user model.User, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"email": email,
	})

	err = u.db.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return user, err
	}

	return user, nil
}

// FindByID :nodoc:
func (u *userRepository) FindByID(id string, user *model.User) error {
	logger := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	err := u.db.Where("id = ?", id).First(user).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// FindAll :nodoc:
func (u *userRepository) FindAll(req map[string]string, userIDs []string) (user []model.User, total int64, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"request": utils.Dump(req),
	})

	qb := u.db.Table("users").Not("id IN (?)", userIDs)

	err = qb.Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err.Error())
		return user, 0, err
	}

	offset, size := utils.GetPageAndSize(req)
	qb.Order("created_at DESC").Offset(offset).Limit(size)

	err = qb.Preload("Photos").Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err.Error())
		return user, 0, err
	}

	return user, total, nil
}
