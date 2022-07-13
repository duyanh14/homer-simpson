package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisHelper struct {
	client *redis.Client
}

func (h *redisHelper) GetString(ctx context.Context, key string, inputValue string) (outValue string, err error) {
	return outValue, nil
}

func (h *redisHelper) SetString(ctx context.Context, key string, inputValue string, time time.Duration) (err error) {
	return nil
}

func (h *redisHelper) Del(ctx context.Context, key string) (err error) {
	return nil
}
