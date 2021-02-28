package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to database")
	}

	fmt.Println("DB Connected: ", db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}