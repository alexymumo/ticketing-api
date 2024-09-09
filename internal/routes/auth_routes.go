package routes

import (
	"events/internal/controllers"
	"events/internal/repository"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	repo := repository.NewAuthRepository(database.Connect())
	route.POST("v1/auth/register", controllers.CreateUser(repo))
	route.POST("v1/auth/login", controllers.LoginUser(repo))
	route.POST("v1/auth/signup", controllers.Register(database.Connect()))
	route.POST("v1/auth/signin", controllers.SignIn(database.Connect()))
	route.GET("v1/auth/test", controllers.Ping())
}
