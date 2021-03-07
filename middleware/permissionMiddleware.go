package middleware

import (
	"admin/database"
	"admin/models"
	"admin/util"
	"errors"
	"github.com/gofiber/fiber"
	"strconv"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")

	Id, err := util.ParseJWT(cookie)

	if err != nil {
		return err
	}

	userId, _ := strconv.Atoi(Id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	database.DB.Preload("Permissions").Find(&role)

	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit_" {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("Unauthorized")
}
