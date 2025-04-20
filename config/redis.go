package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	env := Env()
	redisAddr := fmt.Sprintf("%s:%s", env.Redis.Host, env.Redis.Port)
	redis := redis.NewClient(&redis.Options{
		Addr: redisAddr, Username: env.Redis.User, Password: env.Redis.Pass,
	})
	return redis
}
