package app

import "github.com/tn-go-course/go-search/hw-18/shorturl/pkg/db"

type App struct {
	db db.DB
}

func New(db db.DB) *App {
	return &App{db}
}
