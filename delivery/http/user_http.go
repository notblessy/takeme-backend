package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"
)

// findAllUserHandler :nodoc:
func (h *HTTPService) findAllUserHandler(c echo.Context) error {
	logger := logrus.WithField("context", utils.Dump(c))

	jwtClaim, err := utils.GetSessionClaims(c)
	if err != nil {
		logger.Error(err)
		return utils.Response(c, http.StatusUnauthorized, &utils.HTTPResponse{
			Message: err.Error(),
		})
	}

	request := map[string]string{
		"size": c.QueryParam("size"),
		"page": c.QueryParam("page"),
	}

	users, total, err := h.userUsecase.FindAll(request, jwtClaim.ID)
	if err != nil {
		logger.Error(err)
		return utils.Response(c, http.StatusInternalServerError, &utils.HTTPResponse{
			Message: err.Error(),
		})
	}

	return utils.Response(c, http.StatusOK, &utils.HTTPResponse{
		Data: utils.WithPaging(users, total, c.QueryParam("page"), c.QueryParam("size")),
	})
}
