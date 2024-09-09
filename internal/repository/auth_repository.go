package repository

import (
	"database/sql"
	"events/internal/models"
	"events/pkg/utils"
)

var db *sql.DB

type AuthRepository interface {
	Register(user *models.User) error
	Login(email string) (*models.User, error)
}

type authrepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authrepository{db: db}
}

func (repo *authrepository) Register(user *models.User) error {
	hashpassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashpassword
	stmt, err := db.Prepare("INSERT INTO user (fullname,email,password) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := db.Exec(user.FullName, user.Email, hashpassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.UserID = int(id)
	return err
}

func (repo *authrepository) Login(email string) (*models.User, error) {
	var err error
	return nil, err
	//stmt, err := db.Prepare("SELECT * FROM user WHERE email = ?")
}
