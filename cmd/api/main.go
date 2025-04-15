package main

import (
	"net/http"
	"super-cms/config"
	"super-cms/internal/entity"
	"super-cms/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gorm := config.InitDB()
	err := gorm.AutoMigrate(&entity.Article{})
	if err != nil {
		panic("migrate error")
	}
	route.SetupRoute(gorm, r)
	r.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })
	r.Run()
}
