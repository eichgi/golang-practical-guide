package main

import (
	"admin/database"
	"admin/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
