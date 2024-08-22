package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env values")
	}
	/*
			var (
				user     = os.Getenv("DB_USER")
				password = os.Getenv("DB_PASSWORD")
				host     = os.Getenv("DB_HOST")
				port     = os.Getenv("DB_PORT")
				name     = os.Getenv("DB_NAME")
			)

		var (
			user     = "root"
			password = "1234=Ten"
			host     = "127.0.0.1"
			port     = "3306"
			name     = "ticketdb"
		)
	*/

	//dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "roor", "1234=Ten", "127.0.0.1", "3306", "ticketdb")
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "1234",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "ticketdb",
		AllowNativePasswords: true,
	}
	d, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("connected successfully")
	} else {
		fmt.Println("failed to connect")
	}
	return d
}
