package routes

import (
	"events/internal/controllers"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	route.POST("v1/auth/signup", controllers.Register(database.Connect()))
	route.POST("v1/auth/signin", controllers.SignIn(database.Connect()))
	route.GET("v1/auth/ping", controllers.Ping())
}
