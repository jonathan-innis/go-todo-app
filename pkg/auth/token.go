package auth

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// Create the JWT key used to create the signature
// TODO: Create a key here that is pulled in from some file
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

func ValidateToken(userId string, token string) error {
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return err
	}
	if !parsedToken.Valid {
		return InvalidTokenErr{}
	}
	if claims.UserId != userId {
		return UnauthorizedTokenErr{}
	}
	return nil
}

func ParseToken(authHeader string) string {
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) <= 1 {
		return ""
	}
	return splitToken[1]
}
