package cache

import "context"

type CacheStorage interface {
	SetString(ctx context.Context, key string, value string) error
	GetString(ctx context.Context, key string) (string, error)
}
