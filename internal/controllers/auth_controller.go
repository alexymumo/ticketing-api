package controllers

import (
	"events/internal/models"
	"events/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(repo repository.AuthRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := repo.SignUp(&user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": "successfully created"})
	}
}

func LoginUser(repo repository.AuthRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
