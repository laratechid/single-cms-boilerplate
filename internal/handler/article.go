package handler

import (
	"strconv"
	"super-cms/helper"
	"super-cms/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleHandler struct {
	articleSvc service.ArticleService
}

func NewArticleRoute(db *gorm.DB) ArticleHandler {
	service := service.NewArticleService(db)
	controller := ArticleHandler{
		articleSvc: service,
	}
	return controller
}

func (h ArticleHandler) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.LogErr(err)
		helper.ResErr(c, 400, err.Error())
		return
	}
	data, err := h.articleSvc.GetByID(int64(id))
	if err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	helper.ResSuccess(c, data)
}
