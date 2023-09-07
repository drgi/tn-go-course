package app

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/url"

	"github.com/tn-go-course/go-search/hw-19/shorturl/pkg/db"
)

func (app *App) CreateShortLink(ctx context.Context, rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	id, err := generateId(4)
	if err != nil {
		return "", err
	}
	newUrl := &db.Url{
		Id:  id,
		Url: u.String(),
	}

	err = app.db.StoreUrl(newUrl)
	if err != nil {
		return "", err
	}
	app.SendStat(ctx, newUrl)
	return id, nil
}

func (app *App) RestoreLink(ctx context.Context, id string) (string, error) {
	url, err := app.db.RestoreUrl(id)
	if err != nil {
		return "", err
	}
	return url.Url, nil
}

func generateId(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
