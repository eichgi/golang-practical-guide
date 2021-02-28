package controllers

import (
	"admin/models"
	"github.com/gofiber/fiber"
)

func Register(c *fiber.Ctx) error {

	user := models.User{
		FirstName: "John",
	}

	user.LastName = "Doe"

	//return c.SendString("Hello, World 👋!")
	return c.JSON(user)
}