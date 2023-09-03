package main

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/tn-go-course/lynks/memcache/config"
	"github.com/tn-go-course/lynks/memcache/internal/app"
	"github.com/tn-go-course/lynks/memcache/internal/srv/api"
	"github.com/tn-go-course/lynks/memcache/pkg/cache/redis"
)

const (
	port = 80
)

func main() {
	logger := zerolog.New(os.Stdout)
	logger.Info().Str("App", "running").Msg("running")
	cfg, err := config.Init()
	if err != nil {
		logger.Fatal().Err(err).Msg("init config failed")
		return
	}
	logger = logger.With().
		Str("appName", cfg.App.Name).
		Str("version", cfg.App.Version).Logger()
	redis := redis.New(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.ValueLifeTimeHour)
	app := app.New(&logger, redis)
	api := api.New(app, mux.NewRouter(), &logger)
	api.RegisterHandlers()

	logger.Info().Str("host", cfg.Api.Host).Int("port", cfg.Api.Port).Msg("App listen")
	err = api.Listen(fmt.Sprintf(":%d", cfg.Api.Port))
	if err != nil {
		logger.Fatal().Err(err).Msg("Listen failed")
	}
}
