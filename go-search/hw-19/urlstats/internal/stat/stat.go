package stat

import (
	"fmt"
	"sync"
)

// простые счетчики статистики, не стал усложнять
type UrlCreateStat struct {
	sync.Mutex
	count     int
	avgLength int
}

func (stat *UrlCreateStat) Add(url string) {
	stat.Lock()
	defer stat.Unlock()
	stat.avgLength = (stat.avgLength*stat.count + len(url)) / (stat.count + 1)
	stat.count += 1
}

func (stat *UrlCreateStat) String() string {
	stat.Lock()
	defer stat.Unlock()
	return fmt.Sprintf("Ср. Длинна: %d. Всего: %d", stat.avgLength, stat.count)
}
