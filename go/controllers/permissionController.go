package controllers

import (
	"admin/go/database"
	"admin/go/models"
	"github.com/gofiber/fiber"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}

func CreatePermission(c *fiber.Ctx) error  {
	var permission models.Permission

	if err := c.BodyParser(&permission); err != nil {
		return err
	}

	database.DB.Create(&permission)

	return c.JSON(permission)
}