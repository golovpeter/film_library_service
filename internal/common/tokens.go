package common

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID   int64
	Username string
}

func GenerateJWT(jwtKey string, userID int64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			IssuedAt: &jwt.NumericDate{Time: time.Now()},
		},
		userID,
		username,
	})

	return token.SignedString([]byte(jwtKey))
}

func GetTokenClaims(inputToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(inputToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	return claims, nil
}
