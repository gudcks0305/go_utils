package module

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// Import the missing package for AuthController

// ControllerDI is a struct that holds the fiber app
func AuthControllerDI(app *fiber.App) *AuthController {
	wire.Build(NewAuthController, app)
	return &AuthController{}
}
