package routes

import (
	"events/internal/controllers"
	"events/internal/middlewares"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	route := gin.Default()
	auth := route.Group("/api/v1")
	{
		auth.POST("/signup", controllers.Register(database.Connect()))
		auth.POST("/signin", controllers.SignIn(database.Connect()))
		auth.GET("/users", controllers.GetUsers(database.Connect()))
		auth.DELETE("/user/:userid", controllers.DeleteUser(database.Connect()))
		auth.PUT("/user/:userid", controllers.UpdateUser(database.Connect()))
		auth.GET("/ping", controllers.Ping())
	}

	event := route.Group("/api/v1", middlewares.AuthMiddleware())
	{
		event.POST("/event", controllers.CreateEvent(database.Connect()))
		event.GET("/events", controllers.GetEvents(database.Connect()))
		event.DELETE("/event/:eventid", controllers.DeleteEvent(database.Connect()))
		event.PUT("/event/:eventid", controllers.UpdateEvent(database.Connect()))
		event.GET("/test", controllers.Test())
	}

	ticket := route.Group("/api/v1", middlewares.AuthMiddleware())
	{
		ticket.GET("/pong", controllers.Pong())
		ticket.POST("/create", controllers.CreateTicket(database.Connect()))
		ticket.DELETE("/:ticketid", controllers.CancelTicket())
		ticket.PUT("/:ticketid")
		ticket.GET("/available/:eventid", controllers.AvailableTickets(database.Connect()))
	}

	payment := route.Group("/api/v1", middlewares.AuthMiddleware())
	{
		payment.POST("/pay")
	}
	return route
}
