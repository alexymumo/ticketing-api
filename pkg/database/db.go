package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/ticketdb")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		return nil
	}
	fmt.Println("Connected to db")
	return db
}
