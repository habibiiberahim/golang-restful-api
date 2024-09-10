package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang-restful-api/internal/delivery/http/controller"
	"golang-restful-api/internal/delivery/http/middleware"
	"golang-restful-api/internal/delivery/http/route"
	"golang-restful-api/internal/repository"
	"golang-restful-api/internal/usecase"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	//setup repositories
	userRepository := repository.NewUserRepository(config.Log)

	//setup use case
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, config.Config)

	//setup controller
	userController := controller.NewUserController(userUseCase, config.Log)

	//setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Setup()
}
