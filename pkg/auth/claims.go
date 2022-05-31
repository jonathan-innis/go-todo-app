package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var defaultExpirationTime = time.Hour * 24

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func NewClaims(userId string) *Claims {
	return &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(defaultExpirationTime)),
		},
	}
}
