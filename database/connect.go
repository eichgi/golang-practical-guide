package database

import (
	"admin/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	db, err := gorm.Open(mysql.Open("root:root@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to database")
	}

	fmt.Println("Database connected!")

	db.AutoMigrate(&models.User{})
}
