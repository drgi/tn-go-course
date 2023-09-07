package app

import (
	"context"
	"encoding/json"
	"log"

	"github.com/tn-go-course/go-search/hw-19/urlstats/internal/stat"
)

type url struct {
	Id  string
	Url string
}

type App struct {
	stat *stat.UrlCreateStat
}

func New(s *stat.UrlCreateStat) *App {
	return &App{s}
}

func (a *App) HandleUrlCreate(ctx context.Context, msg []byte) error {
	createdUrl := &url{}
	err := json.Unmarshal(msg, &createdUrl)
	if err != nil {
		return err
	}
	a.stat.Add(createdUrl.Url)
	log.Println(a.stat)
	return nil
}
