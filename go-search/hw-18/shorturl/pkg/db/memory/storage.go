package memory

import (
	"fmt"
	"sync"

	"github.com/tn-go-course/go-search/hw-18/shorturl/pkg/db"
)

type Memory struct {
	sync.Mutex
	list []*db.Url
}

func New() *Memory {
	list := make([]*db.Url, 0)
	return &Memory{list: list}
}

func (m *Memory) StoreUrl(url *db.Url) error {
	m.Lock()
	defer m.Unlock()
	m.list = append(m.list, url)
	return nil
}

func (m *Memory) RestoreUrl(id string) (*db.Url, error) {
	m.Lock()
	defer m.Unlock()
	for _, url := range m.list {
		if url.Id == id {
			return url, nil
		}
	}
	return nil, fmt.Errorf("url not found")
}
