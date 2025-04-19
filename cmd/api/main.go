package main

import (
	"fmt"
	"net/http"
	"super-cms/config"
	"super-cms/docs"
	"super-cms/internal/entity"
	"super-cms/internal/route"
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

// @title						SuperCMS documentation API
// @version					3.0
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {

	// Constructor Dependencies
	r := gin.Default()
	gorm := config.InitDB()
	err := gorm.AutoMigrate(&entity.Article{})
	if err != nil {
		panic("migrate error")
	}

	// Swagger Conf
	docs.SwaggerInfo.BasePath = "/"
	r.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })
	r.GET("/docs/supercms/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Route Setup
	route.SetupRoute(gorm, r)

	r.Run(fmt.Sprintf(":%s", config.Env().App.Port))
}
