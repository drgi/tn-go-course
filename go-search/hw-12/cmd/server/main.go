package main

import (
	"fmt"
	"log"

	"github.com/tn-go-course/go-search/hw-12/pkg/crawler"
	"github.com/tn-go-course/go-search/hw-12/pkg/crawler/spider"
	"github.com/tn-go-course/go-search/hw-12/pkg/index"
	"github.com/tn-go-course/go-search/hw-12/pkg/webapp"
)

const (
	depth = 2
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
	for _, url := range targetUrls {
		docs, err := sp.Scan(url, depth)
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		idx.Create(docs)
	}

	srv := webapp.New(idx)
	srv.RegisterHandlers()
	err := srv.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
