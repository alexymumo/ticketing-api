package controllers

import (
	"events/internal/models"
	"events/internal/repository"
	"events/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Working"})
	}
}
func CreateUser(repo repository.AuthRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
			return
		}
		if err := repo.Register(&user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, user)
	}
}

func LoginUser(repo repository.AuthRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginInput LoginInput

		var user models.User
		if err := ctx.ShouldBindJSON(&loginInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, token, err := repo.Login(user.Email, user.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !utils.VerifyPassword(loginInput.Password, user.Password) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Logged in", "user": result, "token": token})
	}
}
