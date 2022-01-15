package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sepuka/myza/domain"
	"time"
)

type (
	Redis struct {
		store *redis.Client
	}
)

func NewRedis(
	store *redis.Client,
) *Redis {
	return &Redis{
		store: store,
	}
}

func (c *Redis) Close() error {
	return c.store.Close()
}

func (c *Redis) Get(ctx context.Context, key string) domain.CacheResponse {
	return c.store.Get(ctx, key)
}

func (c *Redis) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) domain.CacheResponse {
	return c.store.Set(ctx, key, val, ttl)
}

func (c *Redis) Context() context.Context {
	return c.store.Context()
}
