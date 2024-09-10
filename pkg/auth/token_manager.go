package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang-restful-api/internal/entity"
	"time"
)

type Claims struct {
	Email string
	Name  string
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.User, expirationTime time.Time, secret string) (string, error) {
	byteSecret := []byte(secret)

	claims := &Claims{
		Email: user.Email,
		Name:  user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(byteSecret)
	if err != nil {
		return "JWT Failed for signing token", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string, secret string) (*Claims, error) {
	byteSecret := []byte(secret)
	claims := new(Claims)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return byteSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
