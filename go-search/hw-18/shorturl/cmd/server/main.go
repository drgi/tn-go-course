package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/tn-go-course/go-search/hw-18/shorturl/internal/api"
	"github.com/tn-go-course/go-search/hw-18/shorturl/internal/app"
	"github.com/tn-go-course/go-search/hw-18/shorturl/pkg/db/memory"
)

const (
	port = 80
)

func main() {
	db := memory.New()
	app := app.New(db)
	api := api.New(app, mux.NewRouter())
	api.RegisterHandlers()
	err := api.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
