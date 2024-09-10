package routes

import (
	"events/internal/controllers"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	route.POST("v1/auth/signup", controllers.Register(database.Connect()))
	route.POST("v1/auth/signin", controllers.SignIn(database.Connect()))
	route.GET("v1/auth/users", controllers.GetUsers(database.Connect()))
	route.GET("v1/auth/user{id}")
	route.PUT("v1/auth/user")
	route.DELETE("v1/auth/user/:userid", controllers.DeleteUser(database.Connect()))
	route.GET("v1/auth/ping", controllers.Ping())
}
