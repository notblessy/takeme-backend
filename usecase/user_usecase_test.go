package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/model/mock"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/stretchr/testify/assert"
)

// TestUserUsecase_Register :nodoc:
func TestUserUsecase_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)

	u := userUsecase{
		userRepo: userRepo,
	}

	userID := utils.GenerateID()
	now := time.Now()

	user := model.User{
		ID:          userID,
		Name:        "Frederich Blessy",
		Email:       "frederichblessy@gmail.com",
		Password:    "asdzxc",
		Description: "Its me",
		Gender:      1,
		Preference:  2,
		Age:         26,
		CreatedAt:   now,
		UpdatedAt:   &now,
		DeletedAt:   nil,
	}

	emailFind := model.User{
		ID:    "",
		Name:  "",
		Email: "",
	}

	t.Run("Success", func(t *testing.T) {
		userRepo.EXPECT().FindByEmail(user.Email).
			Times(1).
			Return(user, nil)

		userRepo.EXPECT().Create(user).
			Times(1).
			Return(nil)

		res, err := u.userRepo.FindByEmail(user.Email)
		assert.NoError(t, err)
		assert.NotEqual(t, emailFind, res)

		err = u.userRepo.Create(user)
		assert.NoError(t, err)
	})

	t.Run("email already registered", func(t *testing.T) {
		userRepo.EXPECT().FindByEmail(user.Email).
			Times(1).
			Return(user, nil)

		res, err := u.userRepo.FindByEmail(user.Email)
		assert.NoError(t, err)
		assert.Equal(t, user, res)
	})

	t.Run("Error", func(t *testing.T) {
		userRepo.EXPECT().FindByEmail(user.Email).
			Times(1).
			Return(emailFind, errors.New("internal server error"))

		_, err := u.userRepo.FindByEmail(user.Email)
		assert.Error(t, err)
	})
}

// TestUserUsecase_Login :nodoc:
func TestUserUsecase_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)

	u := userUsecase{
		userRepo: userRepo,
	}

	userID := utils.GenerateID()
	now := time.Now()

	user := model.User{
		ID:          userID,
		Name:        "Frederich Blessy",
		Email:       "frederichblessy@gmail.com",
		Password:    "asdzxc",
		Description: "Its me",
		Gender:      1,
		Preference:  2,
		Age:         26,
		CreatedAt:   now,
		UpdatedAt:   &now,
		DeletedAt:   nil,
	}

	req := model.User{
		Email:    "frederichblessy@gmail.com",
		Password: "asdzxc",
	}
	t.Run("Success", func(t *testing.T) {
		userRepo.EXPECT().FindByEmail(user.Email).
			Times(1).
			Return(user, nil)

		user, err := u.userRepo.FindByEmail(user.Email)
		assert.NoError(t, err)

		if !user.IsPasswordCorrect(req) {
			err = model.ErrIncorrectEmailOrPassword
		}

		assert.NoError(t, err)
	})

	t.Run("Error incorrect email or password", func(t *testing.T) {
		userRepo.EXPECT().FindByEmail(user.Email).
			Times(1).
			Return(user, nil)

		user, err := u.userRepo.FindByEmail(user.Email)
		assert.NoError(t, err)

		if !user.IsPasswordCorrect(model.User{}) {
			err = model.ErrIncorrectEmailOrPassword
		}

		assert.Error(t, err)
	})
}
