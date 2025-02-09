package tokens

import (
	"time"
	"errors"

	"github.com/LavaJover/storage-sso-service/sso-service/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var (
	accessTokenClaims jwt.MapClaims
	cfg = config.MustLoad()
	secretAccessKey, secretRefreshKey = cfg.AccessToken.Secret, cfg.RefreshToken.Secret
)

func GenerateTokens(userID uint64) (string, string, error) {

	// Loading config file
	cfg := config.MustLoad()

	// Creating access token
    accessTokenClaims = jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(cfg.AccessToken.TimeDuration).Unix(),
    }
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
    signedAccessToken, err := accessToken.SignedString([]byte(cfg.AccessToken.Secret))
    if err != nil {
        return "", "", err
    }

    // Creating refresh token
    refreshTokenClaims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(cfg.RefreshToken.TimeDuration).Unix(),
    }
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
    signedRefreshToken, err := refreshToken.SignedString([]byte(cfg.RefreshToken.Secret))
    if err != nil {
        return "", "", err
    }

    return signedAccessToken, signedRefreshToken, nil
}

type Claims struct{
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func ParseAccessJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validating JWT signing algo
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong JWT signing algo")
		}
		return []byte(secretAccessKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Validating Claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token is invalid")
}

func ParseRefreshJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validating JWT signing algo
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong JWT signing algo")
		}
		return []byte(secretRefreshKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Validating Claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token is invalid")
}