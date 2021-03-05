package routes

import (
	"admin/controllers"
	"admin/middleware"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)

	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/other", controllers.Other)
}
