package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisTemplate interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Close() error
}
