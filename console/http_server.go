package console

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/notblessy/takeme-backend/cacher"
	"github.com/notblessy/takeme-backend/config"
	"github.com/notblessy/takeme-backend/db"
	"github.com/notblessy/takeme-backend/delivery/http"
	"github.com/notblessy/takeme-backend/repository"
	"github.com/notblessy/takeme-backend/usecase"
	"github.com/notblessy/takeme-backend/utils"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var runHTTPServer = &cobra.Command{
	Use:   "httpsrv",
	Short: "run http server",
	Long:  `This subcommand is for starting the http server`,
	Run:   runHTTP,
}

func init() {
	rootCmd.AddCommand(runHTTPServer)
}

func runHTTP(cmd *cobra.Command, args []string) {
	mysql := db.MysqlConnection()
	defer db.CloseMysql(mysql)

	redis := db.RedisConnectionPool()

	cache := cacher.NewCacher()
	cache.SetRedisConnectionPool(redis)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	userRepo := repository.NewUserRepository(mysql)
	notifRepo := repository.NewNotificationRepository(mysql)
	reactionRepo := repository.NewReactionRepository(mysql, cache)
	subscriptionRepo := repository.NewSubscriptionRepository(mysql)

	userUsecase := usecase.NewUserUsecase(userRepo, reactionRepo)
	reactionUsecase := usecase.NewReactionUsecase(reactionRepo, notifRepo, subscriptionRepo)
	subscriptionUsecase := usecase.NewSubscriptionUsecase(subscriptionRepo)

	httpSvc := http.NewHTTPService()
	httpSvc.RegisterUserUsecase(userUsecase)
	httpSvc.RegisterReactionUsecase(reactionUsecase)
	httpSvc.RegisterSubscriptionUsecase(subscriptionUsecase)

	httpSvc.Routes(e)

	logrus.Fatal(e.Start(":" + config.HTTPPort()))
}
