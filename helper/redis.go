package helper

import (
	"fmt"
	"super-cms/constant"
)

func RedisGetUserPrefix(userID int64) string {
	return fmt.Sprintf("%s:%d:", constant.RedisUserPrefix, userID)
}
