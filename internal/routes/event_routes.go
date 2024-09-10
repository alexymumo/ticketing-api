package routes

import (
	"events/internal/controllers"
	"events/pkg/database"

	"github.com/gin-gonic/gin"
)

func EventRoutes(route *gin.Engine) {
	route.POST("v1/event/create", controllers.CreateEvent(database.Connect()))
	route.DELETE("v1/event/delete", controllers.DeleteEvent())
	route.GET("v1/event/events", controllers.GetEvents())   //
	route.PUT("v1/event/update", controllers.UpdateEvent()) //owner
	route.GET("v1/event/test", controllers.Test())
}
