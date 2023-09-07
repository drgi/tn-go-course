package main

import (
	"github.com/tn-go-course/go-search/hw-19/shorturl/pkg/msgbroker/kfk"
	"github.com/tn-go-course/go-search/hw-19/urlstats/internal/app"
	"github.com/tn-go-course/go-search/hw-19/urlstats/internal/stat"
)

const (
	kafkaHost  = "localhost:29092"
	kafkaTopic = "urlCreate"
	kafkaGroup = "urlShorter-stat-handler"
)

func main() {
	k := kfk.New([]string{kafkaHost}, kafkaTopic, kafkaGroup)
	stat := &stat.UrlCreateStat{}
	app := app.New(stat)
	k.Register(kafkaTopic, app.HandleUrlCreate)
	k.Consume()
}
