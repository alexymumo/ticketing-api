package models

import "time"

type User struct {
	UserID    int64      `json:"userid"`
	FullName  string     `json:"fullname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"` // event owner && end users
	CreatedAt *time.Time `json:"created_At"`
}
