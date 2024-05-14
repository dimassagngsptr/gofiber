package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(secretKey, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims:= token.Claims.(jwt.MapClaims)
	claims["email"]=email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	return token.SignedString([]byte(secretKey))
}