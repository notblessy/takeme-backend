package usecase

import (
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userUsecase struct {
	userRepo     model.UserRepository
	reactionRepo model.ReactionRepository
}

// NewUserUsecase :nodoc:
func NewUserUsecase(u model.UserRepository, r model.ReactionRepository) model.UserUsecase {
	return &userUsecase{
		userRepo:     u,
		reactionRepo: r,
	}
}

// Register :nodoc:
func (u *userUsecase) Register(req model.RegisterUser) (model.User, error) {
	logger := logrus.WithFields(logrus.Fields{
		"user": utils.Dump(req),
	})

	existUser, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		logger.Error(err.Error())
		return model.User{}, err
	}

	// validate if exist
	if existUser.IsEmailExists(req.Email) {
		logger.Error(model.ErrEmailAlreadyRegistered.Error())
		return model.User{}, model.ErrEmailAlreadyRegistered
	}

	var user model.User

	user.NewUserFromRequest(req)
	if user.ID == "" {
		user.ID = utils.GenerateID()
	}

	if len(req.Photos) > 0 {
		user.NewUserPhotos(req.Photos)
	}

	err = u.userRepo.Create(user)
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
		logger.Error(model.ErrIncorrectEmailOrPassword.Error())
		return "", model.ErrIncorrectEmailOrPassword
	}

	return resp.ID, nil
}

func (u *userUsecase) FindAll(request map[string]string, userID string) ([]model.UserResponse, int64, error) {
	logger := logrus.WithField("request", utils.Dump(request))

	var user model.User
	err := u.userRepo.FindByID(userID, &user)
	if err != nil {
		logger.Error(err)
		return []model.UserResponse{}, 0, nil
	}

	swipedIDs, err := u.reactionRepo.FindAllSwiped(user.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return []model.UserResponse{}, 0, nil
	}

	swipedIDs = append(swipedIDs, user.ID)

	users, total, err := u.userRepo.FindAll(request, swipedIDs)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return []model.UserResponse{}, 0, nil
	}

	var response []model.UserResponse

	for _, user := range users {
		response = append(response, user.ToResponse())
	}

	return response, total, nil
}
