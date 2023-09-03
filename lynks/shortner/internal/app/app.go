package app

import (
	"github.com/rs/zerolog"
	"github.com/tn-go-course/lynks/shortner/internal/repo"
	"github.com/tn-go-course/lynks/shortner/pkg/cache"
)

type App struct {
	repo   *repo.Repo
	cache  cache.CacheStorage
	logger *zerolog.Logger
}

func New(repo *repo.Repo, cache cache.CacheStorage, logger *zerolog.Logger) *App {
	return &App{
		repo:   repo,
		logger: logger,
		cache:  cache,
	}
}
