package route

import (
	"super-cms/internal/handler"
	"super-cms/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB, gin *gin.Engine) {
	auth := middleware.Authentication()

	articleHandler := handler.NewArticleHandler(db)
	authHandler := handler.NewAuthHandler()

	// Article Route
	articleRouter := gin.Group("/articles")
	articleRouter.GET("/:id", auth, articleHandler.GetByID)
	articleRouter.GET("", auth, articleHandler.GetAll)
	articleRouter.POST("", auth, articleHandler.Create)
	articleRouter.PATCH("/:id", auth, articleHandler.Update)
	articleRouter.DELETE("/:id", auth, articleHandler.Delete)

	// Auth Route
	authRouter := gin.Group("/auth")
	authRouter.POST("/login", authHandler.Login)

}
