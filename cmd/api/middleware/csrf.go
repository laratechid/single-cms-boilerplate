package middleware

import (
	"crypto/sha256"
	"fmt"
	"go-pustaka-api/config"
	"go-pustaka-api/helper"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var env = config.Env()

type HeaderCSRF struct {
	APIKey      *string `header:"X-Api-Key" binding:"required"`
	TempoToken  *string `header:"X-Tempo-Token" binding:"required"`
	RequestTime *string `header:"X-Request-Time" binding:"required"`
}

func logCsrf(elapsedTime, csrfToken int, strTime string) {
	fmt.Println("Elapsed Time : ", elapsedTime)
	fmt.Println("Check CSRF Expired Duration : ", csrfToken)
	fmt.Println("Unix Request Time : ", strTime)
	fmt.Println("Unix Server Time : ", time.Now().Unix())
}

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrfToken, err := strconv.Atoi(env.Csrf.TokenDuration)
		if err != nil {
			helper.ResUnauthorized(c, "unauthorized csrf")
			return
		}

		var headerCsrf HeaderCSRF
		if err := c.ShouldBindHeader(&headerCsrf); err != nil {
			helper.ResUnauthorized(c, "unauthorized bind csrf")
			return
		}

		apiKey := c.Request.Header.Get("X-Api-Key")
		if apiKey != env.App.ApiKey {
			helper.ResUnauthorized(c, "unauthorized x-api-key")
			return
		}
		tempoToken := c.Request.Header.Get("X-Tempo-Token")
		strTime := c.Request.Header.Get("X-Request-Time")
		requestTime, _ := strconv.Atoi(strTime)
		elapsedTime := int(time.Now().Unix()) - requestTime
		logCsrf(elapsedTime, csrfToken, strTime)
		if elapsedTime > csrfToken {
			helper.ResForbidden(c, "unauthorized elapse csrf")
			return
		}
		encryptionText := apiKey + strTime
		sum := fmt.Sprintf("%x", sha256.Sum256([]byte(encryptionText)))
		if tempoToken == sum {
			c.Next()
		} else {
			helper.ResForbidden(c, "unauthorized csrf")
			return
		}
	}
}
