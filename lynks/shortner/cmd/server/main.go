package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/tn-go-course/lynks/shortner/internal/app"
	"github.com/tn-go-course/lynks/shortner/internal/config"
	"github.com/tn-go-course/lynks/shortner/internal/repo"
	"github.com/tn-go-course/lynks/shortner/internal/srv/api"
	"github.com/tn-go-course/lynks/shortner/pkg/cache/http"
	"github.com/tn-go-course/lynks/shortner/pkg/metric"
	"github.com/tn-go-course/lynks/shortner/pkg/postgres"
)

const (
	port = 80
)

func main() {
	ctx := context.Background()
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(os.Stdout)
	logger.Info().Str("App", "running").Msg("running")
	cfg, err := config.Init()
	if err != nil {
		logger.Fatal().Err(err).Msg("init config failed")
		return
	}
	logger = logger.With().
		Str("appName", cfg.App.Name).
		Str("version", cfg.App.Version).
		Logger()

	db, err := postgres.New(ctx, cfg.Postgresql.Uri())
	if err != nil {
		logger.Fatal().Err(err).Msg("init db failed")
		return
	}

	c := http.New(cfg.Cache.Url)
	repo := repo.New(db, &logger)

	app := app.New(repo, c, &logger)
	metric := metric.New()
	api := api.New(app, mux.NewRouter(), metric, &logger)
	api.RegisterHandlers()

	logger.Info().Str("host", cfg.Api.Host).Int("port", cfg.Api.Port).Msg("App listen")
	err = api.Listen(fmt.Sprintf(":%d", cfg.Api.Port))
	if err != nil {
		logger.Fatal().Err(err).Msg("Listen failed")
	}
}
