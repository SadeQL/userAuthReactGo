package main

import (
	"go_auth/donne"
	"go_auth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	donne.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
