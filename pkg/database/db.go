package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env values")
	}
	var (
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		name     = os.Getenv("DB_NAME")
	)

	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	d, err := sql.Open("mysql", dburl)
	if err != nil {
		fmt.Println("connected successfully")
	} else {
		fmt.Println("failed to connect")
	}
	return d
}
