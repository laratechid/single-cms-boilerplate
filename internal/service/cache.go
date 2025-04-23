package service

import (
	"context"
	"fmt"
	"go-pustaka-api/helper"
	"time"

	"github.com/go-stack/stack"
	"github.com/redis/go-redis/v9"
)

type CacheService interface {
	Set(key string, value any, exp time.Duration) error
	Get(key string) (any, error)
	Exists(key string) (bool, error)
	DelByPattern(ctx context.Context, pattern string) error
}

type cacheService struct {
	client *redis.Client
	ctx    context.Context
}

func NewCacheService(redisClient *redis.Client) CacheService {
	return &cacheService{
		client: redisClient,
		ctx:    context.Background(),
	}
}

func (r *cacheService) traceErr(err error) {
	stack := stack.Caller(1).Frame().Function
	helper.LogErr(err, stack)
}

func (s *cacheService) Set(key string, value any, exp time.Duration) error {
	if err := s.client.Set(s.ctx, key, value, exp).Err(); err != nil {
		s.traceErr(err)
		return err
	}
	return nil

}

func (s *cacheService) Get(key string) (any, error) {
	data, err := s.client.Get(s.ctx, key).Result()
	if err != nil {
		s.traceErr(err)
		return nil, err
	}
	return data, nil
}

func (s *cacheService) Exists(key string) (bool, error) {
	exists, err := s.client.Exists(s.ctx, key).Result()
	if err != nil {
		s.traceErr(err)
		return false, err
	}
	return exists > 0, err
}

func (s *cacheService) DelByPattern(ctx context.Context, pattern string) error {
	var cursor uint64
	var keysToDelete []string

	for {
		var keys []string
		var err error

		// SCAN for keys matching the pattern
		keys, cursor, err = s.client.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			keysToDelete = append(keysToDelete, keys...)
		}

		if cursor == 0 {
			break
		}
	}

	if len(keysToDelete) > 0 {
		deleted, err := s.client.Del(ctx, keysToDelete...).Result()
		if err != nil {
			return err
		}
		helper.LogInfo(fmt.Sprintf("Deleted %d keys\n", deleted))
	} else {
		helper.LogInfo("No keys found for deletion")
	}

	return nil
}
