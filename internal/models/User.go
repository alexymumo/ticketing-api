package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID    int    `gorm:"primary_key" json:"userid"`
	FullName  string `gorm:"not null" json:"fullname"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `gorm:"not null" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
