package index

import (
	"sort"
	"strings"

	"github.com/tn-go-course/go-search/hw-3/pkg/crawler"
	"github.com/tn-go-course/go-search/hw-3/pkg/str"
)

type Index struct {
	index map[string][]int
	list  []crawler.Document
}

// Создает или добавляет документы в индекс и массив документов
func (i *Index) Create(docs []crawler.Document) *Index {
	if i.index == nil {
		i.index = make(map[string][]int)
	}
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
