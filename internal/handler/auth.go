package handler

import (
	"super-cms/config"
	"super-cms/helper"
	"super-cms/internal/dto"
	"super-cms/internal/repository"
	"super-cms/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authSvc service.AuthService
}

func NewAuthHandler(db *gorm.DB, redis *redis.Client) AuthHandler {
	userRepo := repository.NewUserRepository(db)
	service := service.NewAuthervice(userRepo, redis)
	controller := AuthHandler{
		authSvc: service,
	}
	return controller
}

//	@Tags		Auth
//	@Summary	Auth Login
//	@Param		request	body		dto.AuthRequestDto	true	"Auth payload"
//	@Success	200		{object}	helper.Response{data=string}
//	@Router		/auth/login [post]
func (h AuthHandler) Login(c *gin.Context) {
	var payload dto.AuthRequestDto
	if err := c.Bind(&payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	if err := helper.ValidateRequest(payload); err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	token, err := h.authSvc.Login(c, payload)
	if err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	// temporary
	c.SetCookie("access_token", *token, 3600, "/", config.Env().BaseUrl.FrontendUrl, false, true)
	helper.ResSuccess(c, token)
}

//	@Tags		Auth
//	@Summary	Auth Info
//	@Success	200	{object}	helper.Response{data=dto.UserInfoResponse}
//	@Router		/auth/info [get]
func (h AuthHandler) UserInfo(c *gin.Context) {
	token, err := c.Cookie("access_token")
	if err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	data, err := h.authSvc.UserInfo(token)
	if err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	helper.ResSuccess(c, data)
}
