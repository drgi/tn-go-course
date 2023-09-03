package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/tn-go-course/lynks/shortner/internal/app"
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
	a.router.HandleFunc("/{id}", a.Redirect).Methods(http.MethodGet)
	a.router.HandleFunc("/short", a.GetShortURL).Methods(http.MethodPost)
	a.router.HandleFunc("/url/{id}", a.GetURL).Methods(http.MethodGet)
}

func (s *Api) Listen(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
