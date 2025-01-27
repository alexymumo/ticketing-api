package routes

import (
	"events/internal/controllers"
	"events/internal/middlewares"
	"events/pkg/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func AuthRoutes(route *gin.Engine) {
	auth := route.Group("/api/v1", middlewares.RateLimiter(redisClient, 10, time.Minute))
	{
		auth.POST("/signup", controllers.Register(database.Connect()))
		auth.POST("/login", controllers.SignIn(database.Connect()))
		auth.GET("/ping", controllers.Ping())
	}
	protected := route.Group("/api/v1", middlewares.AuthMiddleware(), middlewares.RateLimiter(redisClient, 10, time.Minute))
	{
		protected.DELETE("/user/:userid", controllers.DeleteUser(database.Connect()))
		protected.PUT("/user/:userid", controllers.UpdateUser(database.Connect()))
		protected.GET("/users", controllers.GetUsers(database.Connect()))
	}
}
