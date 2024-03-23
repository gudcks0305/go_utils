package config

import (
	"github.com/gofiber/fiber/v2"
)

type FiberRegisterRoute struct {
	app *fiber.App
}

func NewFiberRegisterRoute(app *fiber.App) *FiberRegisterRoute {
	return &FiberRegisterRoute{app: app}
}

func (frr *FiberRegisterRoute) RegisterRoutes(controller FiberController) {
	controller.RegisterRoutes(frr.app)
}
