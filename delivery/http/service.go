package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/notblessy/takeme-backend/model"
)

// HTTPService :nodoc:
type HTTPService struct {
	userUsecase     model.UserUsecase
	reactionUsecase model.ReactionUsecase
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// RegisterUserUsecase :nodoc:
func (h *HTTPService) RegisterUserUsecase(u model.UserUsecase) {
	h.userUsecase = u
}

// RegisterReactionUsecase :nodoc:
func (h *HTTPService) RegisterReactionUsecase(r model.ReactionUsecase) {
	h.reactionUsecase = r
}

// Routes :nodoc:
func (h *HTTPService) Routes(route *echo.Echo) {
	v1 := route.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/login", h.loginHandler)
	auth.POST("/register", h.registerUserHandler)

	// Protected routes
	routes := v1.Group("")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middleware.CORS())
	// routes.Use(echojwt.WithConfig(utils.JwtConfig()))

	reaction := routes.Group("/reaction")
	reaction.POST("/", h.createReactionHandler)

}
