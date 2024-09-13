package controllers

import (
	"database/sql"
	"events/internal/models"
	"events/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
		hashedpassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return
		}
		_, err = db.Exec("INSERT INTO user(fullname,email,password) VALUES (?,?,?)", user.FullName, user.Email, hashedpassword)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to register user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "registered successfully",
			"user":    user,
		})
	}
}

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userid := ctx.Param("userid")
		_, err := db.Exec("DELETE FROM user WHERE userid = ?", userid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "successfuly deleted"})

	}
}

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userid := ctx.Param("userid")
		var user models.User
		_, err := db.Exec("UPDATE user SET fullname = ?,email = ?,password = ? WHERE userid = ?", user.FullName, user.Email, user.Password, userid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
			return
		}
		ctx.JSON(http.StatusOK, "updated successfully")
	}
}

func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []models.User
		row, err := db.Query("SELECT userid, fullname, email, password FROM user")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query db"})
			return
		}
		defer row.Close()
		for row.Next() {
			var user models.User
			if err := row.Scan(&user.UserID, &user.FullName, &user.Email, &user.Password); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query db"})
				return
			}
			users = append(users, user)
		}
		if err := row.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get data"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"users": users})
	}
}

func SignIn(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginInput LoginInput
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
		checkPassword := utils.VerifyPassword(storedpassword, loginInput.Password)
		if !checkPassword {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}
		token, err := utils.GenerateToken(loginInput.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generated token"})
			return
		}
		/*
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
		*/

		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}
