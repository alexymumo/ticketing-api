package database

import (
	"events/internal/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:1234@tcp(127.0.0.1:3306)/ticketdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connected successfully")
	} else {
		fmt.Println("failed to connect")
	}
	db.AutoMigrate(&models.User{})
	return db
}
