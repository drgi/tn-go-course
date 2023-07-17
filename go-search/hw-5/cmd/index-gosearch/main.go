package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tn-go-course/go-search/hw-5/pkg/crawler"
	"github.com/tn-go-course/go-search/hw-5/pkg/crawler/spider"
	"github.com/tn-go-course/go-search/hw-5/pkg/index"
)

const (
	depth         = 2
	cacheFileName = "cache.json"
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

	flag.StringVar(&query, "s", "", "input word for search")
	flag.Parse()

	sp = spider.New()
	idx := index.New()
	err := restoreFromFile(idx, cacheFileName)
	if err != nil {
		fmt.Println("WARN:", "cache not restored", err)
	}

	if err != nil {
		for _, url := range targetUrls {
			docs, err := sp.Scan(url, depth)
			if err != nil {
				fmt.Println("ERROR", err)
				continue
			}
			idx.Create(docs)
		}
	}

	err = storeToFile(idx, cacheFileName)
	if err != nil {
		fmt.Println("WARN:", "cache not stored", err)
	}

	docIds := idx.FindIndexes(query)
	docs := make([]crawler.Document, 0, len(docIds))
	for _, docId := range docIds {
		docs = append(docs, idx.FindDoc(docId))
	}
	printResult(docs)

}

// добавил две вспомогательные ф-ии restoreFromFile, storeToFile
//  1. Из двойной обработки ошибки, мне кажеться удобнее в main просто получить
//     конечный результат, получилось восстановить или нет. Флаг добавил для читаемости if в main
//  2. Возможно им место в пакете index, но мне показалось нет, так как нашему index дела до этого нет
//     ему достаточно Reader и Writer)
func restoreFromFile(idx *index.Index, fileName string) error {
	r, err := os.Open(fileName)
	if err != nil {
		return err
	}
	err = idx.Restore(r)
	if err != nil {
		return err
	}
	return nil
}

func storeToFile(idx *index.Index, fileName string) error {
	r, err := os.Create(fileName)
	if err != nil {
		return err
	}
	return idx.Store(r)
}

// Функция вывода на экран
func printResult(docs []crawler.Document) {
	for i, doc := range docs {
		fmt.Printf("Ссылка № %d: %s\n", i+1, doc.URL)
	}

	fmt.Printf("Всего документов найдено: %d.", len(docs))
}
