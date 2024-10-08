package main

import (
	"events/internal/routes"
	"events/pkg/database"
)

func main() {
	database.Connect()
	r := routes.Routes()

	r.Run(":8000")
}
