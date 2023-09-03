package app

import (
	"github.com/rs/zerolog"
	"github.com/tn-go-course/lynks/memcache/pkg/cache"
)

type App struct {
	logger *zerolog.Logger
	redis  cache.CacheStorage
}

func New(logger *zerolog.Logger, redis cache.CacheStorage) *App {
	return &App{
		redis:  redis,
		logger: logger,
	}
}
