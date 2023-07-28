package http

import (
	"github.com/labstack/echo/v4"
	"github.com/notblessy/takeme-backend/model"
)

// HTTPService :nodoc:
type HTTPService struct {
	userUsecase model.UserUsecase
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// RegisterUserUsecase :nodoc:
func (h *HTTPService) RegisterUserUsecase(u model.UserUsecase) {
	h.userUsecase = u
}

// Routes :nodoc:
func (h *HTTPService) Routes(route *echo.Echo) {
	route.POST("/sign-in", h.signInHandler)
	route.POST("/register", h.registerUserHandler)
}
