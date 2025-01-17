package routes

import (
	"events/internal/controllers"
	"events/internal/middlewares"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func AuthRoutes() *gin.Engine {
	route := gin.Default()
	auth := route.Group("/api/v1")
	{
		auth.POST("/signup", controllers.Register(database.Connect()))
		auth.POST("/login", controllers.SignIn(database.Connect()))
		auth.GET("/ping", controllers.Ping())
	}
	r := route.Group("/api/v1", middlewares.AuthMiddleware())
	{
		r.DELETE("/user/:userid", controllers.DeleteUser(database.Connect()))
		r.PUT("/user/:userid", controllers.UpdateUser(database.Connect()))
		r.GET("/users", controllers.GetUsers(database.Connect()))
	}
	return route
}
