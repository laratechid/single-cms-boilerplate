package helper

import (
	"fmt"
	"go-pustaka-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name,omitempty"`
	Username string   `json:"username,omitempty"`
	Email    string   `json:"email,omitempty"`
	Permits  []string `json:"permits,omitempty"`
	Foto     string   `json:"foto,omitempty"`
	Role     string   `json:"role,omitempty"`
	jwt.RegisteredClaims
}

var secretKey = []byte(config.Env().Jwt.Secret)

func GenerateJwtToken(payload JwtPayload) (*string, error) {
	payload.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
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
		LogErr(err)
		return err
	}

	if _, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		return nil
	}
	return err
}

func ParseToken(token string) (*JwtPayload, error) {
	var payload JwtPayload
	if _, _, err := jwt.NewParser().ParseUnverified(token, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
