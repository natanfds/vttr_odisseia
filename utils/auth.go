package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwKey = []byte("secret")

func GenerateJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
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
