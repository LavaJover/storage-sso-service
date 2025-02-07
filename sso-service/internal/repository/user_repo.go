package repo

import (
	"github.com/LavaJover/storage-sso-service/sso-service/pkg/models"
	"gorm.io/gorm"
)

type UserRepo struct{
	*gorm.DB
}

// Insert new user instance to DB
func (repo *UserRepo) CreateUser (newUser *models.User) error{
	result := repo.DB.Create(&newUser)

	if result.Error != nil{
		return result.Error
	}

	return nil
}