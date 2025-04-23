package middleware

import (
	"fmt"
	"go-pustaka-api/helper"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

func AuthWithPermits(redis *redis.Client, permitName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			helper.ResUnauthorized(c, "access key required")
			return
		}
		if err := helper.VerifyJwtToken(accessToken); err != nil {
			helper.ResForbidden(c, err.Error())
			return
		}
		claims, _ := helper.ParseToken(accessToken)
		key := fmt.Sprintf("%s%s", helper.RedisGetUserPrefix(claims.ID), accessToken)
		token, err := redis.Get(c, key).Result()
		if err != nil {
			helper.ResUnauthorized(c, "session was end")
			return
		}
		if err := helper.VerifyJwtToken(token); err != nil {
			helper.ResForbidden(c, err.Error())
			return
		}
		payload, _ := helper.ParseToken(token)
		hasPermission := lo.Contains(payload.Permits, permitName)
		if !hasPermission {
			helper.ResForbidden(c, "no permission access")
			return
		} else {
			c.Next()
		}
		c.Next()
	}
}
