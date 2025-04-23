package route

import (
	mw "go-pustaka-api/cmd/api/middleware"
	"go-pustaka-api/internal/handler_cms"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewsroomRoute(db *gorm.DB, r *gin.Engine, redis *redis.Client) {
	newsroom := r.Group("/newsroom")
	articleHandler := handler_cms.NewArticleHandler(db)
	authHandler := handler_cms.NewAuthHandler(db, redis)
	permits := func(permit string) gin.HandlerFunc {
		return mw.AuthWithPermits(redis, permit)
	}

	// Auth Route
	authRouter := newsroom.Group("/auth")
	authRouter.POST("/login", authHandler.Login)
	authRouter.GET("/info", authHandler.UserInfo)

	// Article Route
	articleRouter := newsroom.Group("/articles")
	articleRouter.GET("/:id", permits("VIEW_ARTICLE_DETAIL"), articleHandler.GetByID)
	articleRouter.GET("", permits("VIEW_ARTICLE_LIST"), articleHandler.GetAll)
	articleRouter.POST("", permits("CREATE_ARTICLE"), articleHandler.Create)
	articleRouter.PATCH("/:id", permits("UPDATE_ARTICLE"), articleHandler.Update)
	articleRouter.DELETE("/:id", permits("DELETE_ARTICLE"), articleHandler.Delete)
}
