package servie

import (
	repo "github.com/LavaJover/storage-sso-service/sso-service/internal/repository"
	"github.com/LavaJover/storage-sso-service/sso-service/pkg/models"
)

type SSOService struct{
	*repo.UserRepo
}

func (service *SSOService) CreateUser (newUser *models.User) error{
	return service.UserRepo.CreateUser(newUser)
}