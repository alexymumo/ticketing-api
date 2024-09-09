package controllers

import (
	"database/sql"
	"events/internal/models"
	"events/internal/repository"
	"events/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var jwtSecret = []byte("ETEAAREAAFD1212")

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Pong"})
	}
}

func Register(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to hash password"})
			return
		}
		_, err = db.Exec("INSERT INTO user(fullname,email,password) VALUES (?,?,?)", user.FullName, user.Email, hashedpassword)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to register user"})
			return
		}
		ctx.JSON(200, gin.H{"message": "registered successfully"})
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
		users, err := repo.Login(user.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !utils.VerifyPassword(users.Password, loginInput.Password) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := utils.GenerateToken(user.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"user": token})
	}
}
func SignIn(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginInput LoginInput
		//var user models.User
		if err := ctx.ShouldBindJSON(&loginInput); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var storedpassword string
		err := db.QueryRow("SELECT password FROM user WHERE email = ?", loginInput.Email).Scan(&storedpassword)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(400, gin.H{"error": "invalid email or password"})
			} else {
				ctx.JSON(500, gin.H{"error": "failed to get user"})
			}
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(storedpassword), []byte(loginInput.Password))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid password"})
			return
		}
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &utils.Claims{
			Email: loginInput.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to generate token"})
			return
		}

		ctx.JSON(200, gin.H{"token": tokenString})
	}
}
