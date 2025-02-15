package controllers

import (
	"database/sql"
	"events/internal/models"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func GetEvents(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var events []models.Event
		result, err := db.Query("SELECT eventid,title,imageUrl,date,venue,description,time,amount,capacity FROM event")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query db"})
			return
		}
		defer db.Close()
		for result.Next() {
			var event models.Event
			err := result.Scan(&event.EventID, &event.Title, &event.ImageUrl, &event.Date, &event.Venue, &event.Description, &event.Time, &event.Amount, &event.Capacity)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query db"})
				return
			}
			events = append(events, event)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"events": events,
		})
	}
}

func CreateEvent(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var event models.Event
		if err := ctx.ShouldBind(&event); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
			return
		}
		cloudinaryurl := "cloudinary://491247926192336:rDATHkoej_xmx0gFS8Ynl69OmLI@doqmtkirc"
		cld, err := cloudinary.NewFromURL(cloudinaryurl)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to init cloudinary"})
		}
		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "ticket"})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload"})
			return
		}
		event.ImageUrl = result.SecureURL
		res, err := db.Exec("INSERT INTO event (title,imageurl,date,venue,description,time,amount,capacity) VALUES (?,?,?,?,?,?,?,?)", event.Title, event.ImageUrl, event.Date, event.Venue, event.Description, event.Time, event.Amount, event.Capacity)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create event"})
			return
		}
		eventID, err := res.LastInsertId()
		if err != nil {
			return
		}
		event.EventID = eventID

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Successfully Created",
			"event":   event,
		})
	}
}

func UpdateEvent(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var event models.Event
		eventid := ctx.Param("eventid")
		_, err := db.Exec("UPDATE event SET title = ? WHERE eventid = ?", event.Title, eventid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "failed to update event")
			return
		}
		ctx.JSON(http.StatusOK, "updated successfully")
	}
}

func DeleteEvent(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventid := ctx.Param("eventid")
		result, err := db.Exec("DELETE FROM event WHERE eventid = ?", eventid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "failed to delete event")
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no affected rows"})
			return
		}
		if rowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
			return
		}
		ctx.JSON(http.StatusOK, "deleted event successfully")
	}
}

func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "test"})
	}
}
