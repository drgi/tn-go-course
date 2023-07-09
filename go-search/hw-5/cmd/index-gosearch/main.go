package main

import (
	"fmt"

	"github.com/tn-go-course/go-search/hw-5/pkg/crawler"
)

const (
	depth = 2
)

var (
	targetUrls = []string{
		"https://go.dev",
		"https://golang.org",
	}
)

func main() {
}

// Функция вывода на экран
func printResult(docs []crawler.Document) {
	for i, doc := range docs {
		fmt.Printf("Ссылка № %d: %s\n", i+1, doc.URL)
	}

	fmt.Printf("Всего документов найдено: %d.", len(docs))
}
