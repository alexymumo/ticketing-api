package controllers

import (
	"context"
	"database/sql"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func UploadToCloudinary(file *multipart.FileHeader) (string, error) {
	cloudinaryurl := "cloudinary://491247926192336:rDATHkoej_xmx0gFS8Ynl69OmLI@doqmtkirc"

	cld, err := cloudinary.NewFromURL(cloudinaryurl)

	var ctx = context.Background()
	result, err := cld.Upload.Upload(ctx, "assets/upload/"+file.Filename, uploader.UploadParams{Folder: "ticket"})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return result.SecureURL, nil

}
func CreateEvent(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cloudinaryurl := "cloudinary://491247926192336:rDATHkoej_xmx0gFS8Ynl69OmLI@doqmtkirc"
		cld, err := cloudinary.NewFromURL(cloudinaryurl)
		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		/*
			err = ctx.SaveUploadedFile(file, "assets/uploads/"+file.Filename)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
				return
			}
		*/

		result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "ticket"})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Successfully uploaded",
			"imageUrl": result.SecureURL,
		})
		/*
			var event models.Event
			if err := ctx.ShouldBind(&event); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
				return
			}
			file, _ := ctx.FormFile("image")
			fileheader, _ := file.Open()

				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": "no image found"})
					return
				}
		*/
		//fmt.Printf("uploaded : %+v\n", header.Filename)
		/*tempPath := "temp_" + file.Filename
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
		c := context.Background()
		result, err := cld.Upload.Upload(c, fileheader, uploader.UploadParams{
			Folder: "ticket",
		})
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
			"message":  "event created successfully",
			"event":    event,
			"imageurl": event.ImageUrl,
		})
		*/
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
