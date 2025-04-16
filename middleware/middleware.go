package middleware

import (
	"strings"
	"super-cms/helper"

	"github.com/gin-gonic/gin"
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
