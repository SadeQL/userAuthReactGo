package donne

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:Quentin81417!@/usersDB"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
}
