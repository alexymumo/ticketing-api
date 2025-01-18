package routes

import (
	"events/internal/controllers"
	"events/internal/middlewares"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(route *gin.Engine) {
	ticket := route.Group("/api/v1", middlewares.AuthMiddleware())
	{
		ticket.GET("/pong", controllers.Pong())
		ticket.POST("/create", controllers.CreateTicket(database.Connect()))
		ticket.DELETE("/:ticketid", controllers.CancelTicket())
		ticket.PUT("/:ticketid")
		ticket.GET("/available/:eventid", controllers.AvailableTickets(database.Connect()))
	}
}
