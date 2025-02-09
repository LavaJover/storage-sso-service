package service

import (
	"fmt"
	"time"

	repo "github.com/LavaJover/storage-sso-service/sso-service/internal/repository"
	tokens "github.com/LavaJover/storage-sso-service/sso-service/internal/token"
	"github.com/LavaJover/storage-sso-service/sso-service/pkg/models"
)

type SSOService struct{
	*repo.UserRepo
}

func (service *SSOService) CreateUser (newUser *models.User) error{
	return service.UserRepo.CreateUser(newUser)
}

func (service *SSOService) GetUserByEmail (email string) (*models.User, error){
	return service.UserRepo.GetUserByEmail(email)
}

/*
	If validation is positive - returns userID
	Otherwise - error with description 
*/

func (service *SSOService) ValidateJWT (tokenString string) (uint64, error) {
	
	if tokenString == ""{
		return 0, fmt.Errorf("token is empty")
	}

	claims, err := tokens.ParseJWT(tokenString)
	if err != nil {
		return 0, fmt.Errorf("token validation error: %v", err)
	}

	// Validating token exp time
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return 0, fmt.Errorf("token has expired")
	}

	return claims.UserID, nil

}