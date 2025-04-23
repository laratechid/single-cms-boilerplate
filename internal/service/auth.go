package service

import (
	"context"
	"fmt"
	"go-pustaka-api/constant"
	"go-pustaka-api/helper"
	"go-pustaka-api/internal/dto"
	"go-pustaka-api/internal/repository"

	"github.com/go-stack/stack"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
)

type AuthService interface {
	Login(ctx context.Context, dto dto.AuthRequestDto) (*string, error)
	UserInfo(token string) (*dto.UserInfoResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
	redis    CacheService
}

func NewAuthervice(userRepo repository.UserRepository, redis *redis.Client) AuthService {
	return &authService{
		userRepo: userRepo,
		redis:    NewCacheService(redis),
	}
}

func (r *authService) traceErr(err error) {
	stack := stack.Caller(1).Frame().Function
	helper.LogErr(err, stack)
}

func (s *authService) Login(ctx context.Context, dto dto.AuthRequestDto) (*string, error) {
	user, err := s.userRepo.GetByEmail(dto.Email)
	if err != nil {
		s.traceErr(err)
		return nil, err
	}
	payload := helper.JwtPayload{ID: user.ID}
	if err = helper.CompareHashPassword(user.Password, dto.Password); err != nil {
		s.traceErr(err)
		return nil, err
	}
	token, err := helper.GenerateJwtToken(payload)
	if err != nil {
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
	tokenWithPermits, err := helper.GenerateJwtToken(payload)
	if err != nil {
		s.traceErr(err)
		return nil, err
	}
	prefixKey := helper.RedisGetUserPrefix(user.ID)
	if err = s.redis.DelByPattern(ctx, fmt.Sprintf("%s%s", prefixKey, "*")); err != nil {
		s.traceErr(err)
		return nil, err
	}
	userKey := fmt.Sprintf("%s%s", prefixKey, *token)
	if err = s.redis.Set(userKey, tokenWithPermits, constant.RedisUserExpiration); err != nil {
		s.traceErr(err)
		return nil, err
	}
	return token, nil
}

func (s *authService) UserInfo(token string) (*dto.UserInfoResponse, error) {
	payload, err := helper.ParseJwtToken(token)
	if err != nil {
		s.traceErr(err)
		return nil, err
	}
	prefixKey := helper.RedisGetUserPrefix(payload.ID)
	userKey := fmt.Sprintf("%s%s", prefixKey, token)
	data, err := s.redis.Get(userKey)
	if err != nil {
		s.traceErr(err)
		return nil, err
	}
	payload, err = helper.ParseJwtToken(data.(string))
	if err != nil {
		s.traceErr(err)
		return nil, err
	}
	result := dto.UserInfoResponse{}
	copier.Copy(&result, &payload)
	return &result, nil
}
