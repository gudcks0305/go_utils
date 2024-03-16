package contoller

import (
	"github.com/gofiber/fiber/v2"
)

// AuthController is a struct that holds the fiber app
type AuthController struct {
	App *fiber.App
}

// NewAuthController is a function that returns a new AuthController

func NewAuthController(app *fiber.App) *AuthController {
	auth := ac.App.Group("/api/v1/auth")
	auth.Post("/login", ac.Login)
	return &AuthController{App: app}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("Login")
	}(c)
}
