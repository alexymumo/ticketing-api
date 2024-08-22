package controllers

import (
	"events/internal/models"
	"events/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Working"})
	}
}
func CreateUser(repo repository.AuthRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		//var register RegisterInput
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
			return
		}
		user, err := repo.Register(user)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, user)
	}
}

func LoginUser(repo repository.AuthRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginInput LoginInput
		//var user models.User
		if err := ctx.ShouldBindJSON(&loginInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := repo.Login(loginInput.Email, loginInput.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}
