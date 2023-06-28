package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/tn-go-course/go-search/pkg/crawler/spider"
	"github.com/tn-go-course/go-search/pkg/flags"
)

const (
	searchQueryFlag = "s"
)

var (
	targetUrls = []string{
		"https://go.dev",
		"https://golang.org",
	}
)

func main() {
	target, err := flags.RequireFlag(searchQueryFlag)
	if err != nil {
		log.Fatal(err)
	}
	app := New(spider.New())
	result, err := app.Search(target)
	if err != nil {
		fmt.Println(err)
	}
	printResult(result)
}

func printResult(r *SearchResult) {
	fmt.Println(strings.Join(r.Urls, "\n"))
	fmt.Printf("Всего документов найдено: %d.\tСовпадение найдено в %d документах.", r.FoundAll, r.Filtred)
}
