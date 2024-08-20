package routes

import (
	"events/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	route.POST("v1/auth/register", controllers.CreateUser)
	route.POST("v1/auth/login", controllers.LoginUser)
}
