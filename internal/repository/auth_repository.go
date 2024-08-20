package repository

import (
	"database/sql"
	"events/internal/models"
	"events/pkg/database"
)

var db *sql.DB

func init() {
	database.Connect()
}

type AuthRepository interface {
	SignUp(user *models.User) error
	SignIn()
}

func Register(user *models.User) error {
	stmt, err := db.Prepare("INSERT INTO users(fullname,email,password, role) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := db.Exec(user.FullName, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.UserID = id
	return err
}

func Login() {

}
