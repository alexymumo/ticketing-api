package models

type User struct {
	UserID   int    `json:"userid"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
