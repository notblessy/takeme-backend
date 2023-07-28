package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

// loginHandler :nodoc:
func (h *HTTPService) loginHandler(c echo.Context) error {
	var req model.AuthRequest

	if err := c.Bind(&req); err != nil {
		return utils.Response(c, http.StatusBadRequest, &utils.HTTPResponse{
			Message: fmt.Sprintf("error bind request: %s", model.ErrBadRequest.Error()),
		})
	}

	if err := c.Validate(&req); err != nil {
		return utils.Response(c, http.StatusBadRequest, &utils.HTTPResponse{
			Message: fmt.Sprintf("error validate: %s", model.ErrBadRequest.Error()),
		})
	}

	user := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(c),
		"request": utils.Dump(user),
	})

	userID, err := h.userUsecase.Login(user)
	if err != nil {
		logger.Error(err)
		switch err {
		case gorm.ErrRecordNotFound:
			return utils.Response(c, http.StatusOK, &utils.HTTPResponse{
				Message: model.ErrBadRequest.Error(),
			})
		case model.ErrIncorrectEmailOrPassword:
			return utils.Response(c, http.StatusBadRequest, &utils.HTTPResponse{
				Message: model.ErrBadRequest.Error(),
			})
		default:
			return utils.Response(c, http.StatusInternalServerError, &utils.HTTPResponse{
				Message: err.Error(),
			})
		}
	}

	token, err := utils.GenerateJwtToken(userID)
	if err != nil {
		return utils.Response(c, http.StatusUnauthorized, &utils.HTTPResponse{
			Message: fmt.Sprintf("token error: %s", model.ErrUnauthorized.Error()),
		})
	}

	return utils.Response(c, http.StatusOK, &utils.HTTPResponse{
		Data: model.Auth{
			ID:    user.ID,
			Token: token,
		},
	})
}

// registerUserHandler :nodoc:
func (h *HTTPService) registerUserHandler(c echo.Context) error {
	logger := logrus.WithField("context", utils.Dump(c))
	var req model.User

	if err := c.Bind(&req); err != nil {
		logger.Error(err)
		return utils.Response(c, http.StatusBadRequest, &utils.HTTPResponse{
			Message: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		logger.Error(err)
		return utils.Response(c, http.StatusBadRequest, &utils.HTTPResponse{
			Message: fmt.Sprintf("error validate: %s", err),
		})
	}

	user, err := h.userUsecase.Register(req)
	if err != nil {
		logger.Error(err)
		return utils.Response(c, http.StatusInternalServerError, &utils.HTTPResponse{
			Message: err.Error(),
		})
	}

	return utils.Response(c, http.StatusCreated, &utils.HTTPResponse{
		Data: user,
	})
}
