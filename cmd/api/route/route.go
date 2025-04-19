package route

import (
	mw "super-cms/cmd/api/middleware"
	"super-cms/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB, gin *gin.Engine) {
	auth := mw.Authentication()
	articleHandler := handler.NewArticleHandler(db)
	authHandler := handler.NewAuthHandler(db)

	// Auth Route
	authRouter := gin.Group("/auth")
	authRouter.POST("/login", authHandler.Login)

	// Article Route
	articleRouter := gin.Group("/articles")
	articleRouter.GET("/:id", auth, mw.Permit("VIEW_ARTICLE_DETAIL"), articleHandler.GetByID)
	articleRouter.GET("", auth, mw.Permit("VIEW_ARTICLE_LIST"), articleHandler.GetAll)
	articleRouter.POST("", auth, mw.Permit("CREATE_ARTICLE"), articleHandler.Create)
	articleRouter.PATCH("/:id", auth, mw.Permit("UPDATE_ARTICLE"), articleHandler.Update)
	articleRouter.DELETE("/:id", auth, mw.Permit("DELETE_ARTICLE"), articleHandler.Delete)
}
