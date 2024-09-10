package controllers

import (
	"context"
	"database/sql"
	"events/internal/models"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func initcloudinary() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams("", "", "")
	if err != nil {
		return nil, err
	}
	return cld, err
}
func CreateEvent(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var event models.Event
		if err := ctx.ShouldBindJSON(&event); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no image found"})
			return
		}
		tempPath := "temp_" + file.Filename
		if err := ctx.SaveUploadedFile(file, tempPath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
			return
		}
		defer os.Remove(tempPath)
		cld, err := initcloudinary()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cloudinary init failed"})
			return
		}
		var context = context.Background()
		result, err := cld.Upload.Upload(context, tempPath, uploader.UploadParams{Folder: "tickets"})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload"})
			return
		}
		event.ImageUrl = result.SecureURL

		_, err = db.Exec("INSERT INTO event (title,imageurl,date,venue,description,time,amount,capacity) VALUES (?,?,?,?,?,?,?,?)", event.Title, event.ImageUrl, event.Date, event.Venue, event.Description, event.Time, event.Amount, event.Capacity)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to create event"})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "event created successfully",
			"event":   event,
		})
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

func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "test"})
	}
}
