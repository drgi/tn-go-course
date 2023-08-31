package main

import (
	"fmt"

	"github.com/tn-go-course/go-search/hw-11/pkg/crawler"
	"github.com/tn-go-course/go-search/hw-11/pkg/crawler/spider"
	"github.com/tn-go-course/go-search/hw-11/pkg/index"
	"github.com/tn-go-course/go-search/hw-11/pkg/netsrv"
)

const (
	depth = 2
	port  = 8000
)

var (
	targetUrls = []string{
		"https://go.dev",
		"https://golang.org",
	}
)

func main() {
	var sp crawler.Interface
	sp = spider.New()
	idx := index.New()
	fmt.Println("Wait scan...")
	for _, url := range targetUrls {
		docs, err := sp.Scan(url, depth)
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		idx.Create(docs)
	}
	s := netsrv.New(10)
	s.RegisterHandler(idx.NetHandler)
	fmt.Printf("Listen at: %d", port)
	err := s.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}
}
