package config

import "github.com/gofiber/fiber/v2"

type FiberController interface {
	RegisterRoutes(app *fiber.App)
}

// Path: web/fiber/controller_interface.go
// Compare this snippet from script/process_killer_port/kill_process.go:
