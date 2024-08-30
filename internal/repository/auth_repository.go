package repository

import (
	"errors"
	"events/internal/models"
	"events/pkg/utils"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, string, error)
}

type authrepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authrepository{db: db}
}

func (repo *authrepository) Register(user *models.User) error {
	hashpassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashpassword
	return repo.db.Create(&user).Error
}

func (repo *authrepository) Login(email, password string) (*models.User, string, error) {
	user := models.User{}
	result := repo.db.Where("email=?", email).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("error occured")
		}
		return nil, "", result.Error
	}
	if !utils.VerifyPassword(password, user.Password) {
		return nil, "", errors.New("Error")
	}
	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return nil, "", errors.New("Error")
	}
	return &user, token, nil
}
