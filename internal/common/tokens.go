package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID   int64
	Username string
	Role     string
}

func GenerateJWT(jwtKey string, userID int64, username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			IssuedAt: &jwt.NumericDate{Time: time.Now()},
		},
		userID,
		username,
		role,
	})

	return token.SignedString([]byte(jwtKey))
}
