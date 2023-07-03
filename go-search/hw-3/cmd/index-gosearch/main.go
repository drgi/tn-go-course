package main

import (
	"flag"
	"fmt"

	"github.com/tn-go-course/go-search/hw-3/pkg/crawler"
	"github.com/tn-go-course/go-search/hw-3/pkg/crawler/spider"
	"github.com/tn-go-course/go-search/hw-3/pkg/index"
)

const (
	searchQueryFlag = "s"
	deth            = 2
)

var (
	targetUrls = []string{
		"https://go.dev",
		"https://golang.org",
	}
)

func main() {
	var documents []crawler.Document
	var sp crawler.Interface
	var query string

	flag.StringVar(&query, searchQueryFlag, "", "input word for search")
	flag.Parse()

	sp = spider.New()

	for _, url := range targetUrls {
		docs, err := sp.Scan(url, deth)
		if err != nil {
			fmt.Println(err)
			continue
		}
		documents = append(documents, docs...)
	}
	idx := &index.Index{}
	idx.Create(documents)
}
