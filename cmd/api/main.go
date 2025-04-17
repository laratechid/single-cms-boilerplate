package main

import (
	"fmt"
	"net/http"
	"super-cms/config"
	"super-cms/docs"
	"super-cms/internal/entity"
	"super-cms/internal/route"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title						SuperCMS documentation API
// @version					    3.0
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })
	r.GET("/docs/supercms/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	gorm := config.InitDB()
	err := gorm.AutoMigrate(&entity.Article{})
	if err != nil {
		panic("migrate error")
	}
	route.SetupRoute(gorm, r)
	r.Run(fmt.Sprintf(":%s", config.Env().App.Port))
}
