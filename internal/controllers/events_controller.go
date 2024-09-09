package controllers

import (
	"database/sql"
	"events/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var event models.Event
		if err := ctx.ShouldBindJSON(&event); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		_, err := db.Exec("INSERT INTO event () VALUES ()")
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to create event"})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "event created successfully"})
	}

}

func UpdateEvent() gin.HandlerFunc {
	return func(ctx *gin.Context) {}

}

func DeleteEvent() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func GetEvents() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
