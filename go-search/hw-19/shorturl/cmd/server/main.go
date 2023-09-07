package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/tn-go-course/go-search/hw-19/shorturl/internal/api"
	"github.com/tn-go-course/go-search/hw-19/shorturl/internal/app"
	"github.com/tn-go-course/go-search/hw-19/shorturl/pkg/db/memory"
	"github.com/tn-go-course/go-search/hw-19/shorturl/pkg/msgbroker/kfk"
)

const (
	port = 80

	kafkaHost  = "localhost:29092"
	kafkaTopic = "urlCreate"
	kafkaGroup = "urlShorter"
)

func main() {
	db := memory.New()

	k := kfk.New([]string{kafkaHost}, kafkaTopic, kafkaGroup)
	app := app.New(db, k)
	api := api.New(app, mux.NewRouter())
	api.RegisterHandlers()
	err := api.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
