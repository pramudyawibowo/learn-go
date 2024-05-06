package helper

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = os.Getenv("JWT_SECRET_KEY")

type JWTClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

func GenerateJWT(id uint, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "issuer",
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": JWTClaims{
			ID:   id,
			Role: role,
		},
	})
	return token.SignedString([]byte(secret))
}

func VerifyJWT(tokenString string) (*JWTClaims, error) {
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

	data := claims["data"].(map[string]interface{})

	return &JWTClaims{
		ID:   uint(data["id"].(float64)),
		Role: data["role"].(string),
	}, nil
}
