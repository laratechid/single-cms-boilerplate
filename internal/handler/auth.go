package handler

import (
	"super-cms/helper"
	"super-cms/internal/dto"
	"super-cms/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authSvc service.AuthService
}

func NewAuthHandler() AuthHandler {
	service := service.NewAuthervice()
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
	token, err := h.authSvc.Login(payload)
	if err != nil {
		helper.ResErr(c, 400, err.Error())
		return
	}
	helper.ResSuccess(c, token)
}
