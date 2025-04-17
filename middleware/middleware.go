package middleware

import (
	"errors"
	"strings"
	"super-cms/config"
	"super-cms/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-stack/stack"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 1 {
			helper.ResUnauthorized(c, "token required")
			return
		}
		authorization := strings.Split(authHeader, " ")
		if len(authorization) <= 1 {
			helper.ResUnauthorized(c, "malformed token")
			return
		}
		jwtToken := authorization[1]
		if err := helper.VerifyJwtToken(jwtToken); err != nil {
			helper.ResForbidden(c, "invalid token")
			return
		}
		c.Next()
	}
}

func ValidatePermission(c *gin.Context) error {
	var secretKey = []byte(config.Env().Jwt.Secret)
	authHeader := c.GetHeader("Authorization")
	authorization := strings.Split(authHeader, " ")
	claims := &helper.JwtPayload{}
	jwt.ParseWithClaims(authorization[1], claims, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	stackName := stack.Caller(1).Frame().Function
	hasPermission := lo.Contains(claims.Permits, stackName)
	if hasPermission {
		return nil
	} else {
		return errors.New("no permission access")
	}
}
