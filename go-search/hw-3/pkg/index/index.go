package index

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/tn-go-course/go-search/hw-3/pkg/crawler"
)

type Index struct {
	index map[string][]int
	list  []crawler.Document
}

func (i *Index) Create(docs []crawler.Document) *Index {
	i.index = make(map[string][]int)
	for _, d := range docs {
		fmt.Printf("Doc: %+v. ID: %v\n", d.Title, i.generateID())
		title := strings.TrimSpace(strings.ReplaceAll(d.Title, ""))
	}
	return i
}

func (i *Index) Find(query string) []int {
	return i.index[query]
}

func (i *Index) generateID() int64 {
	var id int64
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		id += int64(rand.Intn(1000))
	}
	fmt.Println(id)
	ts := time.Now().UnixMilli() + id
	return ts
}
