package http

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/notblessy/takeme-backend/model"
	"github.com/notblessy/takeme-backend/utils"
)

// HTTPService :nodoc:
type HTTPService struct {
	userUsecase         model.UserUsecase
	reactionUsecase     model.ReactionUsecase
	subscriptionUsecase model.SubscriptionUsecase
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

// RegisterSubscriptionUsecase :nodoc:
func (h *HTTPService) RegisterSubscriptionUsecase(s model.SubscriptionUsecase) {
	h.subscriptionUsecase = s
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
	routes.Use(echojwt.WithConfig(utils.JwtConfig()))

	reactions := routes.Group("/reactions")
	reactions.POST("", h.createReactionHandler)

	users := routes.Group("/users")
	users.GET("", h.findAllUserHandler)

	subscriptions := routes.Group("/subscriptions")
	subscriptions.POST("", h.createSubscriptionHandler)

}
