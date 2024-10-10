package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ping"})
	}
}

func CreateTicket() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//eventid := ctx.Param("eventid")
		//userid,exists :=
		//var event models.Event

	}
}
