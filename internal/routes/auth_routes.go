package routes

import (
	"events/internal/controllers"
	"events/internal/middlewares"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	route := gin.Default()
	auth := route.Group("/auth")
	{
		auth.POST("/signup", controllers.Register(database.Connect()))
		auth.POST("/signin", controllers.SignIn(database.Connect()))
		auth.GET("/users", controllers.GetUsers(database.Connect()))
		auth.DELETE("/user/:userid", controllers.DeleteUser(database.Connect()))
		auth.GET("/ping", controllers.Ping())
	}
	// protected routes
	event := route.Group("/event", middlewares.AuthMiddleware())
	{
		event.POST("/event", controllers.CreateEvent(database.Connect()))
		event.GET("/events", controllers.GetEvents())
		event.GET("/test", controllers.Test())
	}
	return route
}
