package main

import (
	"fmt"
	"log"

	"github.com/tn-go-course/go-search/hw-13/internal/api"
	"github.com/tn-go-course/go-search/hw-13/internal/index"
	"github.com/tn-go-course/go-search/hw-13/pkg/crawler"
	"github.com/tn-go-course/go-search/hw-13/pkg/crawler/spider"
)

const (
	depth = 1
	port  = 80
)

var (
	targetUrls = []string{
		"https://go.dev",
		"https://golang.org",
	}
)

func main() {
	var sp crawler.Interface
	sp = spider.New()
	idx := index.New()
	fmt.Println("Wait scan...")
	err := idx.Scan(sp, targetUrls, depth)
	if err != nil {
		log.Fatal(err)
	}

	srv := api.New(idx)
	srv.RegisterHandlers()
	err = srv.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
