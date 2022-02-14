package database

import (
	"go_auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:Quentin81417!@/usersDB"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	connection.AutoMigrate(&models.User{})
}
