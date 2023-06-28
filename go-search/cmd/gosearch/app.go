package main

import (
	"github.com/tn-go-course/go-search/pkg/crawler"
	"github.com/tn-go-course/go-search/pkg/str"
)

// Структура результата поиска
type SearchResult struct {
	Query    string
	FoundAll int
	Filtred  int
	Urls     []string
}

// App - основная структура приложения
type App struct {
	crawler crawler.Interface
}

func New(crawler crawler.Interface) *App {
	return &App{crawler: crawler}
}

// Основной метод поиска
func (a *App) Search(query string) (*SearchResult, error) {
	result := &SearchResult{
		Query: query,
	}
	documents, err := a.scanUrls(targetUrls)
	if err != nil {
		return nil, err
	}
	result.FoundAll = len(documents)
	documents = a.filterDocuments(documents, query)
	result.Filtred = len(documents)
	result.Urls = make([]string, 0, cap(documents))
	for _, doc := range documents {
		result.Urls = append(result.Urls, doc.URL)
	}
	return result, nil
}

/* Внутрение, вспомогательные функции */

// сканирует по массиву юрл
func (s *App) scanUrls(urls []string) ([]crawler.Document, error) {
	documents := make([]crawler.Document, 0)
	for _, url := range urls {
		docs, err := s.crawler.Scan(url, 2)
		if err != nil {
			return documents, err
		}
		documents = append(documents, docs...)
	}
	return documents, nil
}

// Фильтрует массив доментов по параметру, запросу в поиске
func (a *App) filterDocuments(docs []crawler.Document, query string) []crawler.Document {
	filtredDocs := make([]crawler.Document, 0)
	for _, d := range docs {
		if a.containsInDocument(&d, query) {
			filtredDocs = append(filtredDocs, d)
		}
	}
	return filtredDocs
}

// Проверяет заданное слово по полям в структуре документа.
func (a *App) containsInDocument(doc *crawler.Document, query string) bool {
	return str.Contains(doc.Title, query, true) || str.Contains(doc.Body, query, true)
}
