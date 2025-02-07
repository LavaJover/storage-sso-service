package repo

import (
	"github.com/LavaJover/storage-sso-service/sso-service/pkg/models"
	"gorm.io/gorm"
)

type UserRepo struct{
	*gorm.DB
}

func (repo *UserRepo) CreateUser (user *models.User) error{

}