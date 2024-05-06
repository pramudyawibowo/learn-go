package helper

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

var secret = os.Getenv("JWT_SECRET_KEY")

func GenerateJWT(id uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(secret)), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
