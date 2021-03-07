package main

import (
	"admin/go/database"
	"admin/go/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"log"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	//app.Listen(":8000")
	log.Fatal(app.Listen("localhost:8000"), nil)
}
