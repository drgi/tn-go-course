package main

import (
	"flag"
	"fmt"

	"github.com/tn-go-course/go-search/hw-2/pkg/crawler"
	"github.com/tn-go-course/go-search/hw-2/pkg/crawler/spider"
	"github.com/tn-go-course/go-search/hw-2/pkg/str"
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
	documents = filterDocuments(documents, query)
	printResult(documents)
}

// Фильтрует массив доментов по параметру, запросу в поиске
func filterDocuments(docs []crawler.Document, query string) []crawler.Document {
	filtredDocs := make([]crawler.Document, 0)
	for _, d := range docs {
		if str.Contains(d.Title, query, true) || str.Contains(d.Body, query, true) {
			filtredDocs = append(filtredDocs, d)
		}
	}
	return filtredDocs
}

// Функция вывода на экран
func printResult(docs []crawler.Document) {
	for i, doc := range docs {
		fmt.Printf("Ссылка № %d: %s\n", i+1, doc.URL)
	}

	fmt.Printf("Всего документов найдено: %d.", len(docs))
}
