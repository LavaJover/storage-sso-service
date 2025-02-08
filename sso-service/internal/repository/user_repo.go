package repo

import (
	"log/slog"

	"github.com/LavaJover/storage-sso-service/sso-service/pkg/models"
	"gorm.io/gorm"
)

type UserRepo struct{
	*gorm.DB
}

// Insert new user instance to DB
func (repo *UserRepo) CreateUser (newUser *models.User) error{
	result := repo.DB.Create(newUser)

	if result.Error != nil{
		return result.Error
	}

	return nil
}

// Get user from database by email
func (repo *UserRepo) GetUserByEmail (email string) (*models.User, error){
	user := &models.User{}
	result := repo.DB.Where("email = ?", email).First(user)

	if result.Error != nil{
		slog.Error("failed to find user by email", "msg", result.Error.Error())
		return nil, result.Error
	}
	return user, nil
}