package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(secretKey string, payload map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	for key, value := range payload {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func GenerateRefreshToken(secretKey string, payload map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for key, value := range payload {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshTokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}
