package tokens

import (
	"time"

	"github.com/LavaJover/storage-sso-service/sso-service/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(userID uint64) (string, string, error) {

	// Loading config file
	cfg := config.MustLoad()

	// Creating access token
    accessTokenClaims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(cfg.AccessToken.TimeDuration).Unix(),
    }
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
    signedAccessToken, err := accessToken.SignedString(cfg.AccessToken.Secret)
    if err != nil {
        return "", "", err
    }

    // Creating refresh token
    refreshTokenClaims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(cfg.RefreshToken.TimeDuration).Unix(),
    }
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
    signedRefreshToken, err := refreshToken.SignedString(cfg.RefreshToken.Secret)
    if err != nil {
        return "", "", err
    }

    return signedAccessToken, signedRefreshToken, nil
}