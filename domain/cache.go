package domain

import (
	"context"
	"time"
)

type (
	CacheResponse interface {
		Val() string
		Err() error
	}

	Cache interface {
		Get(ctx context.Context, key string) CacheResponse
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) CacheResponse
		Context() context.Context
	}
)
