package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tn-go-course/go-search/hw-13/pkg/crawler"
	"golang.org/x/exp/slog"
)

type App interface {
	Search(ctx context.Context, query string) ([]crawler.Document, error)
	CreateDocument(ctx context.Context, doc crawler.Document) (int, error)
	UpdateDocument(ctx context.Context, doc crawler.Document) error
	DeleteDocument(ctx context.Context, id int) error
}

type Api struct {
	router *mux.Router
	app    App
	logger *slog.Logger
}

func New(app App) *Api {
	router := mux.NewRouter()
	logger := slog.Default()
	return &Api{
		logger: logger,
		router: router,
		app:    app,
	}
}

func (a *Api) RegisterHandlers() {
	a.router.Use(a.setRequestId)
	a.router.Use(a.logRequest)
	a.router.Use(a.setTimeoutAndRecovery)

	a.router.HandleFunc("/api/v1/search", a.Search).Methods(http.MethodGet)
	a.router.HandleFunc("/api/v1/document", a.CreateDocument).Methods(http.MethodPost)
	a.router.HandleFunc("/api/v1/document", a.UpdateDocument).Methods(http.MethodPatch)
	a.router.HandleFunc("/api/v1/document/{id}", a.DeleteDocument).Methods(http.MethodDelete)
}

func (s *Api) Listen(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
