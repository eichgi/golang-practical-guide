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

	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.AllPermissions)
	app.Post("/api/permissions", controllers.CreatePermission)

	app.Get("/api/products", controllers.AllProducts)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Post("/api/products", controllers.CreateProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

	app.Post("/api/upload", controllers.Upload)

	app.Static("/api/uploads", "./uploads")

	app.Get("/other", controllers.Other)
}
