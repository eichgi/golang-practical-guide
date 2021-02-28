package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:root@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to database")
	}

	fmt.Println("Database connected!")
}
