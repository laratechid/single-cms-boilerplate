package route

import (
	"go-pustaka-api/cmd/api/middleware"
	"go-pustaka-api/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ClientRoute(db *gorm.DB, r *gin.Engine, redis *redis.Client) {
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.SecurityMiddleware())
	r.Use(middleware.CSRFMiddleware())
	client := r.Group("/client")
	articleHandler := handler.NewArticleHandler(db)

	// Article Route
	articleRouter := client.Group("/articles")
	articleRouter.GET("/:id", articleHandler.GetByID)
	articleRouter.GET("", articleHandler.GetAll)
}
