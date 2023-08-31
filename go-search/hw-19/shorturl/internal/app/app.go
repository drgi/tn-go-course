package app

import (
	"github.com/tn-go-course/go-search/hw-19/shorturl/pkg/db"
	"github.com/tn-go-course/go-search/hw-19/shorturl/pkg/msgbroker"
)

type App struct {
	db     db.DB
	broker msgbroker.MessageBroker
}

func New(db db.DB, broker msgbroker.MessageBroker) *App {
	return &App{db, broker}
}
