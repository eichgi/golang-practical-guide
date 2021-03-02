package routes

import (
	"admin/controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)
	app.Get("/other", controllers.Other)
}
