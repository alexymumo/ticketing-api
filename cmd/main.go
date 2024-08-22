package main

import (
	"events/internal/routes"
	"events/pkg/database"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env values")
	}
	database.Connect()
	route := gin.Default()
	route.Use(gin.Logger())
	routes.AuthRoutes(route)
	route.Run(":8000")
}
