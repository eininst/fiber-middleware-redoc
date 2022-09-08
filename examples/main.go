package main

import (
	redoc "github.com/eininst/fiber-middleware-redoc"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/redoc/*", redoc.New("examples/swagger.json"))

	app.Listen(":8080")
}
