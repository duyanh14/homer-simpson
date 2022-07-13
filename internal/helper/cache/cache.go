package cache

import (
	"context"
	"simpson/config"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type CacheHelper interface {
	GetString(ctx context.Context, key string, value string) (string, error)
	SetString(ctx context.Context, key string, value string, time time.Duration) error
	Del(ctx context.Context, key string) error
}

func NewRedisInstance(ctx context.Context, cfg config.Redis) (CacheHelper, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addrs[0],
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return &redisHelper{
			client: client,
		}, err
	}
	zap.S().Debug("connect to redis successful")
	return &redisHelper{
		client: client,
	}, nil
}

// InitRedis create a redis from config
func InitRedis(ctx context.Context, cfg config.Redis) (*redis.Client, error) {
	var redisClient *redis.Client
	opts, err := redis.ParseURL(cfg.Addrs[0])
	if err != nil {
		return nil, err
	}
	if cfg.PoolSize != 0 {
		opts.PoolSize = cfg.PoolSize
	}
	if cfg.WriteTimeoutSeconds != 0 {
		opts.WriteTimeout = time.Duration(cfg.WriteTimeoutSeconds) * time.Second
	}
	if cfg.IdleTimeoutSeconds != 0 {
		opts.IdleTimeout = time.Duration(cfg.IdleTimeoutSeconds) * time.Second
	}
	if cfg.ReadTimeoutSeconds != 0 {
		opts.ReadTimeout = time.Duration(cfg.ReadTimeoutSeconds) * time.Second
	}
	if cfg.DialTimeoutSeconds != 0 {
		opts.DialTimeout = time.Duration(cfg.DialTimeoutSeconds) * time.Second
	}

	redisClient = redis.NewClient(opts)

	cmd := redisClient.Ping(ctx)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	zap.S().Debug("connect to redis successful")
	return redisClient, nil
}
