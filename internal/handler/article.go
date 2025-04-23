package handler

import (
	"go-pustaka-api/helper"
	"go-pustaka-api/internal/dto"
	"go-pustaka-api/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleHandler struct {
	articleSvc service.ArticleService
}

func NewArticleHandler(db *gorm.DB) ArticleHandler {
	service := service.NewArticleService(db)
	controller := ArticleHandler{
		articleSvc: service,
	}
	return controller
}

//	@Tags		Article
//	@Security	XApiKey
//	@Security	XTempoToken
//	@Security	XRequestTime
//	@Summary	Get  Article Details
//	@Param		id	path		int	true	"Article ID"
//	@Success	200	{object}	helper.Response{data=dto.ArticleDetailResponse}
//	@Router		/client/articles/{id} [get]
func (h ArticleHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
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

//	@Tags		Article
//	@Summary	Get All Article
//	@Param		request	query		dto.PaginationRequestDto	true	"Query Params"
//	@Success	200		{object}	helper.Response{data=dto.PaginationResponseDtoExample}
//	@Router		/client/articles [get]
func (h ArticleHandler) GetAll(c *gin.Context) {
	var payload dto.PaginationRequestDto
	if err := c.Bind(&payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	if err := helper.ValidateRequest(payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	data, err := h.articleSvc.GetAll(payload)
	if err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	helper.ResSuccess(c, data)
}
