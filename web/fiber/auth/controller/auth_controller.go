package controller

import (
	"github.com/gofiber/fiber/v2"
)

// AuthController is a struct that holds the fiber app
type AuthController struct {
}

// NewAuthController is a function that returns a new AuthController

func (ac *AuthController) NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) RegisterRoutes(app *fiber.App) {
	group := app.Group("/api/v1/auth")
	group.Get("/login", ac.Login)
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("Login")
	}(c)
}
