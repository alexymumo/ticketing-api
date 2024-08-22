package repository

import (
	"database/sql"
	"errors"
	"events/internal/models"
	"events/pkg/utils"
	"log"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email, password string) (string, error)
}

type authrepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authrepository{db: db}
}

func (repo *authrepository) Register(user models.User) (models.User, error) {
	//var users models.User

	query := "INSERT INTO users (fullname,email,password,role) VALUES (?,?,?,?)"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatal("error on the prepare query", err)
		return user, err
	}
	defer stmt.Close()

	result, err := repo.db.Exec(user.FullName, user.Email, user.Password, user.Role)
	if err != nil {
		return user, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}
	user.UserID = int(id)
	return user, err
}

func (repo *authrepository) Login(email, password string) (string, error) {
	stmt, err := repo.db.Prepare("SELECT fullname,email,password FROM users WHERE email = ?")
	var user models.User
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	err = repo.db.QueryRow(email).Scan(user.FullName, user.Email, &user.Password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	err = user.VerifyPassword(user.Password)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil

}
