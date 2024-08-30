package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("ETEAAREAAFD1212")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString("1212AFDAFADFDAF")
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
