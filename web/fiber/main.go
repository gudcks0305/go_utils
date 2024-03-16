package main

import "github.com/gofiber/fiber/v2"

func main() {
	// fiber web server
	app := fiber.New()
	app.Listen(":3000")

}
