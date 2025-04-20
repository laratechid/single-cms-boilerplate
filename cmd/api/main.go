package main

import (
	"fmt"
	"net/http"
	"super-cms/cmd/api/route"
	"super-cms/config"
	"super-cms/docs"
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

//	@title		SuperCMS documentation API
//	@version	3.0
func main() {

	// Constructor Dependencies
	r := gin.Default()
	gorm := config.InitDB()
	redis := config.InitRedis()

	docs.SwaggerInfo.BasePath = "/"
	r.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })
	r.GET("/docs/supercms/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	route.SetupRoute(gorm, r, redis)

	r.Run(fmt.Sprintf(":%s", config.Env().App.Port))
}
