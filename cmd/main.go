package main

import (
	"events/internal/routes"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.Engine) {
	routes.AuthRoutes(route)
	routes.EventRoutes(route)
	routes.TicketRoutes(route)
}

func main() {
	gin.SetMode(gin.DebugMode)
	database.Connect()
	router := gin.Default()
	InitRoutes(router)
	router.Run(":8000")
}
