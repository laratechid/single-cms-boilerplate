package route

import (
	"super-cms/internal/handler"
	"super-cms/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB, gin *gin.Engine) {
	articleHandler := handler.NewArticleHandler(db)
	authHandler := handler.NewAuthHandler()

	// Article Route
	articleRouter := gin.Group("/articles")
	articleRouter.GET("/:id", middleware.Authentication(), articleHandler.GetByID)
	articleRouter.GET("", articleHandler.GetAll)
	articleRouter.POST("", articleHandler.Create)
	articleRouter.PATCH("/:id", articleHandler.Update)
	articleRouter.DELETE("/:id", articleHandler.Delete)

	// Auth Route
	authRouter := gin.Group("/auth")
	authRouter.POST("/login", authHandler.Login)

}
