package repository

import (
	"events/internal/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) error
	Login(email string) (*models.User, error)
}

type authrepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authrepository{db: db}
}

func (repo *authrepository) Register(user models.User) error {
	return repo.db.Create(&user).Error
}

func (repo *authrepository) Login(email string) (*models.User, error) {
	var user models.User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
