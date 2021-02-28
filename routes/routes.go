package routes

import (
	"admin/controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Get("/other", controllers.Other)
}
