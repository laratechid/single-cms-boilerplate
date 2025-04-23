package helper

import (
	"fmt"
	"go-pustaka-api/constant"
)

func RedisGetUserPrefix(userID int64) string {
	return fmt.Sprintf("%s:%d:", constant.RedisUserPrefix, userID)
}
