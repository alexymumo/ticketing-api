package controllers

import (
	"events/internal/models"
	"events/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string
	Password string
}

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
		ctx.JSON(http.StatusCreated, gin.H{"data": "user created"})
	}
}

func LoginUser(repo repository.AuthRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginInput LoginInput
		if err := ctx.ShouldBindJSON(&loginInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//if err := repo.SignIn(&loginInput); err != nil {
		//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		//	return
		//}
		//ctx.JSON(http.StatusOK, gin.H{"data": "logged in"})

	}
}
