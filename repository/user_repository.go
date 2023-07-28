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

	err := u.db.Create(&user).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}

// FindByEmail :nodoc:
func (u *userRepository) FindByEmail(email string) (user model.User, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"email": email,
	})

	err = u.db.Where("email = ?", email).First(&user).Error
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

// FindAllUsersByRole :nodoc:
func (u *userRepository) FindAllUsersByRole(organizationID, role string, user *[]model.User) error {
	logger := logrus.WithFields(logrus.Fields{
		"organization_id": organizationID,
		"role":            role,
	})

	err := u.db.Where("role = ? AND organization_id = ?", role, organizationID).Find(user).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
