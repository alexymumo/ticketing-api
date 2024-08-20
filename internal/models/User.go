package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    int64      `json:"userid"`
	FullName  string     `json:"fullname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"` // event owner && end users
	CreatedAt *time.Time `json:"created_At"`
}

func (user *User) HashPassword(password string) error {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashpassword)
	return nil
}

func (user *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
