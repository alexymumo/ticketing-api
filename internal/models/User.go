package models

type User struct {
	UserID      int    `json:"userid"`
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Password    string `json:"password"`
}
