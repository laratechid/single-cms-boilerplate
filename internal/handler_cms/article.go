package handler_cms

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

//	@Tags		CMS - Article
//	@Security	BearerAuth
//	@Summary	Get Article Details
//	@Param		id	path		int	true	"Article ID"
//	@Success	200	{object}	helper.Response{data=dto.ArticleDetailResponse}
//	@Router		/newsroom/articles/{id} [get]
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

//	@Tags		CMS - Article
//	@Security	BearerAuth
//	@Summary	Get All Article
//	@Param		request	query		dto.PaginationRequestDto	true	"Query Params"
//	@Success	200		{object}	helper.Response{data=dto.PaginationResponseDtoExample}
//	@Router		/newsroom/articles [get]
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

//	@Tags		CMS - Article
//	@Security	BearerAuth
//	@Summary	Create Article
//	@Param		request	body		dto.ArticleCreateRequestDto	true	"Article payload"
//	@Success	200		{object}	helper.Response{data=string}
//	@Router		/newsroom/articles [post]
func (h ArticleHandler) Create(c *gin.Context) {
	var payload dto.ArticleCreateRequestDto
	if err := c.Bind(&payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	if err := helper.ValidateRequest(payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	if err := h.articleSvc.Create(payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	helper.ResSuccess(c, "ok")
}

//	@Tags		CMS - Article
//	@Security	BearerAuth
//	@Summary	Update Article
//	@Param		id		path		int							true	"Article ID"
//	@Param		request	body		dto.ArticleUpdateRequestDto	true	"Article payload"
//	@Success	200		{object}	helper.Response{data=string}
//	@Router		/newsroom/articles/{id} [patch]
func (h ArticleHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	var payload dto.ArticleUpdateRequestDto
	if err := c.Bind(&payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	if err := helper.ValidateRequest(payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	if err := h.articleSvc.Update(int64(id), payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	helper.ResSuccess(c, "ok")
}

//	@Tags		CMS - Article
//	@Security	BearerAuth
//	@Summary	Delete Article
//	@Param		id	path		int	true	"Article ID"
//	@Success	200	{object}	helper.Response{data=string}
//	@Router		/newsroom/articles/{id} [delete]
func (h ArticleHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	if err := h.articleSvc.Delete(int64(id)); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	helper.ResSuccess(c, "ok")
}
