package main

import (
	"events/internal/routes"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	route := gin.Default()
	route.Use(gin.Logger())
	routes.AuthRoutes(route)
	route.Run(":8000")
}
