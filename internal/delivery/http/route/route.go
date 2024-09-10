package route

import "C"
import (
	"github.com/gofiber/fiber/v2"
	"golang-restful-api/internal/delivery/http/controller"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *controller.UserController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/api/users/current", c.UserController.Current)
	c.App.Delete("/api/users/logout", c.UserController.Logout)
	c.App.Patch("/api/users/update", c.UserController.Update)
}
