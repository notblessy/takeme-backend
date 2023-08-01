package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
)

// createSubscriptionHandler :nodoc:
func (h *HTTPService) createSubscriptionHandler(c echo.Context) error {
	logger := logrus.WithField("context", utils.Dump(c))
	var req model.Subscription

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

	jwtClaim, err := utils.GetSessionClaims(c)
	if err != nil {
		logger.Error(err)
		return utils.Response(c, http.StatusUnauthorized, &utils.HTTPResponse{
			Message: err.Error(),
		})
	}

	req.UserID = jwtClaim.ID

	subs, err := h.subscriptionUsecase.Create(req)
	if err != nil {
		logger.Error(err)
		return utils.Response(c, http.StatusInternalServerError, &utils.HTTPResponse{
			Message: err.Error(),
		})
	}

	return utils.Response(c, http.StatusCreated, &utils.HTTPResponse{
		Data: subs.ID,
	})
}
