package route

import (
	"super-cms/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB, gin *gin.Engine) {
	articleRoute := handler.NewArticleRoute(db)

	articleRouter := gin.Group("/articles")
	articleRouter.GET("/:id", articleRoute.GetByID)
	articleRouter.GET("", articleRoute.GetAll)
	articleRouter.POST("", articleRoute.Create)
	articleRouter.PATCH("/:id", articleRoute.Update)
	articleRouter.DELETE("/:id", articleRoute.Delete)
}
