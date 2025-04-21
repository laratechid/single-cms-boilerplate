package service

import (
	"super-cms/helper"
	"super-cms/internal/dto"
	"super-cms/internal/repository"

	"github.com/go-stack/stack"
	"github.com/jinzhu/copier"
)

type AuthService interface {
	Login(dto dto.AuthRequestDto) (*string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthervice(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (r authService) traceErr(err error) {
	stack := stack.Caller(1).Frame().Function
	helper.LogErr(err, stack)
}

func (s *authService) Login(dto dto.AuthRequestDto) (*string, error) {
	payload := helper.JwtPayload{}
	user, err := s.userRepo.GetByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	if err = helper.CompareHashPassword(user.Password, dto.Password); err != nil {
		s.traceErr(err)
		return nil, err
	}
	copier.Copy(&payload, &user)
	// TODO - fetch permits from user role table
	payload.Permits = append(
		payload.Permits,
		"VIEW_ARTICLE_DETAIL",
		"VIEW_ARTICLE_LIST",
	)
	token, err := helper.GenerateJwtToken(payload)
	if err != nil {
		return nil, err
	}
	return token, nil
}
