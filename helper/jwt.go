package helper

import (
	"fmt"
	"super-cms/config"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Permit   []string `json:"permit"`
	jwt.MapClaims
}

var secretKey = []byte(config.Env().Jwt.Secret)

func GenerateJwtToken(payload JwtPayload) (*string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := claims.SignedString(secretKey)
	if err != nil {
		LogErr(err)
		return nil, err
	}
	return &token, nil
}

func VerifyJwtToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &JwtPayload{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		return nil
	}
	return err
}
