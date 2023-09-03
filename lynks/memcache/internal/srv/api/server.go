package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/tn-go-course/lynks/memcache/internal/app"
)

type Api struct {
	app    *app.App
	router *mux.Router
	logger *zerolog.Logger
}

func New(app *app.App, router *mux.Router, logger *zerolog.Logger) *Api {
	l := logger.With().Str("pkg", "api").Logger()
	return &Api{app, router, &l}
}

func (a *Api) RegisterHandlers() {
	a.router.Use(a.setRequestIdAndTs)
	a.router.Use(a.logRequest)
	a.router.Use(a.setTimeoutAndRecovery)
	a.router.HandleFunc("/store/{id}", a.GetURL).Methods(http.MethodGet)
	a.router.HandleFunc("/store", a.StoreURL).Methods(http.MethodPost)
}

func (s *Api) Listen(addr string) error {
	return http.ListenAndServe(addr, s.router)
}