package token

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte("my_super_secret_key_for_access")
var refreshSecret = []byte("my_super_secret_key_for_refresh")

func generateTokens(userID uint64) (string, string, error) {

	// Creating access token (15 mins duration)
    accessTokenClaims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(15 * time.Minute).Unix(),
    }
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
    signedAccessToken, err := accessToken.SignedString(accessSecret)
    if err != nil {
        return "", "", err
    }

    // Creating refresh token (7 days duration)
    refreshTokenClaims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
    }
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
    signedRefreshToken, err := refreshToken.SignedString(refreshSecret)
    if err != nil {
        return "", "", err
    }

    return signedAccessToken, signedRefreshToken, nil
}