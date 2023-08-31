package index

import (
	"errors"

	"github.com/tn-go-course/go-search/hw-13/pkg/crawler"
)

func (i *Index) Scan(sp crawler.Interface, targetUrls []string, depth int) error {
	for _, url := range targetUrls {
		docs, err := sp.Scan(url, depth)
		if err != nil {
			return errors.New("scan failed. err: " + err.Error())
		}
		i.Create(docs)
	}
	return nil
}
