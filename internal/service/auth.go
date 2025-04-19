package service

import (
	"fmt"
	"super-cms/helper"
	"super-cms/internal/dto"

	"github.com/jinzhu/copier"
)

type AuthService interface {
	Login(dto dto.AuthRequestDto) (*string, error)
}

type authService struct{}

func NewAuthervice() AuthService {
	return &authService{}
}

func (s *authService) Login(dto dto.AuthRequestDto) (*string, error) {
	fmt.Println("1")
	payload := helper.JwtPayload{}
	copier.Copy(&payload, &dto)
	payload.Permits = append(payload.Permits,
		"super-cms/internal/handler.ArticleHandler.GetAll", "super-cms/internal/handler.ArticleHandler.GetByID",
	)
	fmt.Println("2")
	token, err := helper.GenerateJwtToken(payload)
	if err != nil {
		return nil, err
	}
	return token, nil
}
