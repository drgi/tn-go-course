package index

import (
	"context"
	"errors"
	"sort"
	"strings"

	"github.com/tn-go-course/go-search/hw-13/pkg/crawler"
	"github.com/tn-go-course/go-search/hw-13/pkg/str"
)

type Index struct {
	index map[string][]int
	list  []crawler.Document
}

func New() *Index {
	i := &Index{}
	i.index = make(map[string][]int)
	return i
}

// Создает или добавляет документы в индекс и массив документов
func (i *Index) Create(docs []crawler.Document) *Index {
	for _, d := range docs {
		title := str.FilterString(d.Title)
		d.ID = i.generateID()
		words := strings.Split(title, " ")
		i.list = append(i.list, d)
		for _, word := range words {
			word = strings.TrimSpace(word)
			if word == "" {
				continue
			}
			if s, ok := i.index[word]; !ok {
				i.index[word] = []int{d.ID}
			} else {
				i.index[word] = append(s, d.ID)
			}
		}
	}
	i.sort()
	return i
}

// Находит уникальные ИД документов
func (i *Index) FindIndexes(query string) []int {
	query = strings.TrimSpace(strings.ToLower(query))
	iDs, ok := i.index[query]
	if !ok {
		return []int{}
	}
	uniques := make([]int, 0, len(iDs))
	unique := make(map[int]bool)
	for _, id := range iDs {
		if _, ok := unique[id]; !ok {
			uniques = append(uniques, id)
			unique[id] = true
		}
	}

	return uniques
}

// Бинарный поиск по массиву документов
func (i *Index) FindDoc(index int) crawler.Document {
	list := i.list
	start, end := 0, len(list)-1
	for start <= end {
		mid := (start + end) / 2

		if list[mid].ID == index {
			return list[mid]
		}
		if list[mid].ID < index {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return crawler.Document{}
}

// Поиск по запросу с использованием индекса
func (i *Index) Search(_ context.Context, query string) ([]crawler.Document, error) {
	if len(i.list) == 0 {
		return nil, errors.New("empty document list")
	}
	docIds := i.FindIndexes(query)
	docs := make([]crawler.Document, 0, len(docIds))
	for _, docId := range docIds {
		doc := i.FindDoc(docId)
		if doc.ID != 0 {
			docs = append(docs, doc)
		}
	}
	return docs, nil
}

func (i *Index) CreateDocument(ctx context.Context, doc crawler.Document) (int, error) {
	doc.ID = i.generateID()
	i.list = append(i.list, doc)
	return doc.ID, nil
}
func (i *Index) UpdateDocument(ctx context.Context, doc crawler.Document) error {
	arrIndex, err := i.docIndex(doc.ID)
	if err != nil {
		return err
	}
	i.list[arrIndex].Title = doc.Title
	return nil
}
func (i *Index) DeleteDocument(ctx context.Context, id int) error {
	arrIndex, err := i.docIndex(id)
	if err != nil {
		return err
	}
	i.list = append(i.list[:arrIndex], i.list[arrIndex+1:]...)
	return nil
}

// поиск индекса в массиве по ид
func (i *Index) docIndex(id int) (int, error) {
	for i, d := range i.list {
		if d.ID == id {
			return i, nil
		}
	}
	return 0, errors.New("document not found")
}

// Выделил в отдельную ф-ю так как пробовал разную логику генерации числовых ИД, оставил простой вариант
func (i *Index) generateID() int {
	return len(i.list) + 1
}

// сортирует массив документов по ИД
func (i *Index) sort() {
	sort.Slice(i.list, func(x, j int) bool {
		return i.list[x].ID < i.list[j].ID
	})
}
