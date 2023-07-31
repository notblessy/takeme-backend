package usecase

import (
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepo model.UserRepository
}

// NewUserUsecase :nodoc:
func NewUserUsecase(u model.UserRepository) model.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

// Register :nodoc:
func (u *userUsecase) Register(req model.RegisterUser) (model.User, error) {
	logger := logrus.WithFields(logrus.Fields{
		"user": utils.Dump(req),
	})

	var user model.User

	user.NewUserFromRequest(req)
	if user.ID == "" {
		user.ID = utils.GenerateID()
	}

	err := u.userRepo.Create(user)
	if err != nil {
		logger.Error(err.Error())
		return model.User{}, err
	}

	return user, nil
}

// Login :nodoc:
func (u *userUsecase) Login(user model.User) (string, error) {
	logger := logrus.WithField("user", utils.Dump(user))

	resp, err := u.userRepo.FindByEmail(user.Email)
	if err != nil {
		logger.Error(err)
		return "", nil
	}

	if !resp.IsPasswordCorrect(user) {
		logger.Error(err)
		return "", err
	}

	return user.ID, nil
}
