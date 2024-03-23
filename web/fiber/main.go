package main

import (
	"fiber/auth/controller"
	"fiber/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// fiber web server
	app := fiber.New()
	router_register := config.NewFiberRegisterRoute(app)
	authController := &controller.AuthController{}

	// 생성된 포인터를 사용하여 NewAuthController 메서드 호출
	router_register.RegisterRoutes(authController.NewAuthController())
	app.Listen(":3000")

}
