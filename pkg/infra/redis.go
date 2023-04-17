package infra

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"go-clean-arch/config"
)

// InitRedis create a redis from config
func InitRedis(redisCfg *config.RedisConfig) (*redis.Client, error) {
	var redisClient *redis.Client

	opts, err := redis.ParseURL(redisCfg.ConnectionURL)
	if err != nil {
		zap.S().Debugf("parse redis url fail: %+v", err)
		return nil, err
	}

	opts.PoolSize = redisCfg.PoolSize
	opts.DialTimeout = time.Duration(redisCfg.DialTimeoutSeconds) * time.Second
	opts.ReadTimeout = time.Duration(redisCfg.ReadTimeoutSeconds) * time.Second
	opts.WriteTimeout = time.Duration(redisCfg.WriteTimeoutSeconds) * time.Second
	opts.IdleTimeout = time.Duration(redisCfg.IdleTimeoutSeconds) * time.Second

	redisClient = redis.NewClient(opts)

	cmd := redisClient.Ping(context.Background())
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	zap.S().Debug("connect to redis successful")
	return redisClient, nil
}
