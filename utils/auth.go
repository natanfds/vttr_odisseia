package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/natanfds/vtt_odisseia/configs"
)

var jwKey = []byte(configs.ENV.JwtSecret)

func GenerateJWT(userId string) (string, error) {
	tokenExpiration := time.Duration(configs.ENV.TokenExpirationDays) * 24 * time.Hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(tokenExpiration).Unix(),
	})

	return token.SignedString(jwKey)
}

func extractJWT(tk *jwt.Token) (interface{}, error) {
	return jwKey, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, extractJWT)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}
