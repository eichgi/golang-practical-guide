package controllers

import "github.com/gofiber/fiber"

func Other(c *fiber.Ctx) error {
	return c.SendString("sup man!")
}
