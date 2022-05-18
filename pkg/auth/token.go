package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

func GetTokenForUserId(userId string) (string, error) {
	claims := NewClaims(userId)
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := tokenWithClaims.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateToken(userId string, token string) (bool, error) {
	return false, nil
}
