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
	var sp crawler.Interface
	var query string

	flag.StringVar(&query, searchQueryFlag, "", "input word for search")
	flag.Parse()

	sp = spider.New()
	idx := &index.Index{}

	for _, url := range targetUrls {
		docs, err := sp.Scan(url, deth)
		if err != nil {
			fmt.Println(err)
			continue
		}
		idx.Create(docs)
	}

	docIds := idx.FindIndexes(query)
	docs := make([]crawler.Document, 0, len(docIds))
	for _, docId := range docIds {
		docs = append(docs, idx.FindDoc(docId))
	}
	printResult(docs)
}

// Функция вывода на экран
func printResult(docs []crawler.Document) {
	for i, doc := range docs {
		fmt.Printf("Ссылка № %d: %s\n", i+1, doc.URL)
	}

	fmt.Printf("Всего документов найдено: %d.", len(docs))
}
