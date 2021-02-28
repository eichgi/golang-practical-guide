package main

import (
	"admin/database"
	"admin/routes"
	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":8000")
}
