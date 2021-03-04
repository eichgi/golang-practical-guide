package controllers

import (
	"admin/database"
	"admin/models"
	"github.com/gofiber/fiber"
)

func AllUsers(c *fiber.Ctx) error {

	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error  {
	var user models.User

	if err:= c.BodyParser(&user); err != nil {
		return err
	}

	//password, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}