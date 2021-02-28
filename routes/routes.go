package routes

import (
	"admin/controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App)  {
	app.Get("/", controllers.Hello)
	app.Get("/other", controllers.Other)
}