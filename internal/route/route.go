package route

import (
	"super-cms/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB, gin *gin.Engine) {
	articleRouter := gin.Group("/articles")
	articleRoute := handler.NewArticleRoute(db)

	articleRouter.GET("/:id", articleRoute.GetOne)
}
