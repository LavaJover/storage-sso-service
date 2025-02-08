package server

import (
	"context"
	"log/slog"

	"github.com/LavaJover/storage-sso-service/sso-service/internal/password"
	"github.com/LavaJover/storage-sso-service/sso-service/internal/service"
	tokens "github.com/LavaJover/storage-sso-service/sso-service/internal/token"
	"github.com/LavaJover/storage-sso-service/sso-service/pkg/models"
	ssopb "github.com/LavaJover/storage-sso-service/sso-service/proto/gen"
)

type SSOServer struct{
	ssopb.UnimplementedAuthServiceServer
	*service.SSOService
}

func (server *SSOServer) Register(ctx context.Context, r *ssopb.RegisterRequest) (*ssopb.AuthResponse, error){

	// Password encryption
	hashedPassword, err := password.HashPassword(r.Password)

	if err != nil{
		slog.Error("failed to hash password", "msg", err.Error())
		return nil, err
	}

	// Inserting user instance to db
	user := models.User{
		Email: r.Email,
		Password: hashedPassword,
	}

	err = server.SSOService.CreateUser(&user)

	if err != nil{
		slog.Error("failed to create new user", "msg", err.Error())
		return nil, err
	}

	// Generating JWT and refresh tokens
	accessToken, refreshToken, err := tokens.GenerateTokens(uint64(user.ID))

	if err != nil{
		slog.Error("failed to generate tokens for user", "msg", err.Error())
		return nil, err
	}

	// Send response
	return &ssopb.AuthResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		UserId: uint64(user.ID),
	}, nil
}

// func (server *SSOServer) Login(ctx context.Context, r *ssopb.LoginRequest) (*ssopb.AuthResponse, error){

// }

// func (server *SSOServer) ValidateToken(ctx context.Context, r *ssopb.ValidateTokenRequest) (*ssopb.ValidateTokenResponse, error){

// }

// func (server *SSOServer) RefreshToken(ctx context.Context, r *ssopb.RefreshTokenRequest) (*ssopb.AuthResponse, error){

// }