package routes

import (
	"events/internal/controllers"
	"events/internal/middlewares"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func EventRoutes(route *gin.Engine) {
	event := route.Group("/api/v1", middlewares.AuthMiddleware())
	{
		event.POST("/event", controllers.CreateEvent(database.Connect()))
		event.GET("/events", controllers.GetEvents(database.Connect()))
		event.DELETE("/event/:eventid", controllers.DeleteEvent(database.Connect()))
		event.PUT("/event/:eventid", controllers.UpdateEvent(database.Connect()))
		event.GET("/test", controllers.Test())
	}
}
