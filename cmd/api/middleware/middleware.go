package middleware

import (
	"errors"
	"strings"
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
			helper.ResForbidden(c, err.Error())
			return
		}
		c.Next()
	}
}

// Strategy: We attach permit name on route then compare permits on jwt payload
func Permit(permitName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		authorization := strings.Split(authHeader, " ")
		claims := &helper.JwtPayload{}
		if _, _, err := jwt.NewParser().ParseUnverified(authorization[1], claims); err != nil {
			stackName := stack.Caller(0).Frame().Function
			helper.LogErr(err, stackName)
			helper.ResInternalServerError(c, "internal server error")
			return
		}
		hasPermission := lo.Contains(claims.Permits, permitName)
		if !hasPermission {
			helper.ResForbidden(c, "no permission access")
			return
		} else {
			c.Next()
		}
	}
}

// Strategy: We validate permission on handler then compare permits on jwt payload
func ValidatePermission(c *gin.Context) error {
	authHeader := c.GetHeader("Authorization")
	authorization := strings.Split(authHeader, " ")
	claims := &helper.JwtPayload{}
	if _, _, err := jwt.NewParser().ParseUnverified(authorization[1], claims); err != nil {
		return err
	}
	stackName := stack.Caller(0).Frame().Function
	hasPermission := lo.Contains(claims.Permits, stackName)
	if hasPermission {
		return nil
	} else {
		msg := "no permission access"
		err := errors.New(msg)
		helper.LogErr(err, stackName)
		return err
	}
}
