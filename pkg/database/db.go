package database

import (
	"events/internal/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {

	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "1234", "127.0.0.1", "3306", "ticketdb")
	d, err := gorm.Open("mysql", dburl)
	if err != nil {
		fmt.Println("connected successfully")
	} else {
		fmt.Println("failed to connect")
	}
	d.AutoMigrate(&models.User{})
	return d
}
