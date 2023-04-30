package database

import (
	"FiberProject/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:33066)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with database")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.Player{})
}
