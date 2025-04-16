package service

import (
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
	payload := helper.JwtPayload{}
	copier.Copy(&payload, &dto)
	token, err := helper.GenerateJwtToken(payload)
	if err != nil {
		return nil, err
	}
	return token, nil
}
