package main

import (
	"fmt"
	"go-pustaka-api/cmd/api/route"
	"go-pustaka-api/config"
	"go-pustaka-api/constant"
	"go-pustaka-api/docs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:     true,
		TimestampFormat: time.RFC3339,
	})
}

// @title						Go Pustaka API Documentation
// @version					3.0
// @securityDefinitions.apikey	XApiKey
// @in							header
// @name						X-Api-Key
// @securityDefinitions.apikey	XTempoToken
// @in							header
// @name						X-Tempo-Token
// @securityDefinitions.apikey	XRequestTime
// @in							header
// @name						X-Request-Time
func main() {
	r := gin.Default()
	gorm := config.InitDB()
	redis := config.InitRedis()

	docs.SwaggerInfo.BasePath = "/"
	r.GET(constant.SwaggerPath, ginSwagger.WrapHandler(swaggerFiles.Handler))
	route.NewsroomRoute(gorm, r, redis)
	route.ClientRoute(gorm, r, redis)
	r.Run(fmt.Sprintf(":%s", config.Env().App.Port))
}
