package main

import (
	"events/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.New()
	route.Use(gin.Logger())
	routes.AuthRoutes(route)
	route.Run("localhost:8080")
}
