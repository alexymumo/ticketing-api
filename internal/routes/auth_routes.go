package routes

import (
	"events/internal/controllers"
	"events/internal/middlewares"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	auth := route.Group("/api/v1")
	{
		auth.POST("/signup", controllers.Register(database.Connect()))
		auth.POST("/login", controllers.SignIn(database.Connect()))
		auth.GET("/ping", controllers.Ping())
	}
	protected := route.Group("/api/v1", middlewares.AuthMiddleware())
	{
		protected.DELETE("/user/:userid", controllers.DeleteUser(database.Connect()))
		protected.PUT("/user/:userid", controllers.UpdateUser(database.Connect()))
		protected.GET("/users", controllers.GetUsers(database.Connect()))
	}
}
