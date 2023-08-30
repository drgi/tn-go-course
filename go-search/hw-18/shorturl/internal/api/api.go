package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tn-go-course/go-search/hw-18/shorturl/internal/app"
)

type Api struct {
	app    *app.App
	router *mux.Router
}

func New(app *app.App, router *mux.Router) *Api {
	return &Api{app, router}
}

func (a *Api) RegisterHandlers() {
	a.router.HandleFunc("/{id}", a.Redirect).Methods(http.MethodGet)
	a.router.HandleFunc("/short", a.GetShortURL).Methods(http.MethodPost)
	a.router.HandleFunc("/url/{id}", a.GetURL).Methods(http.MethodGet)
}

func (s *Api) Listen(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
