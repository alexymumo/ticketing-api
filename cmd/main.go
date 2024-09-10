package main

import (
	"events/internal/routes"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	route := gin.Default()
	routes.AuthRoutes(route)

	/*
		protected := route.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			routes.EventRoutes(route)
		}
	*/
	route.Use(gin.Logger())

	route.Run(":8000")
}
