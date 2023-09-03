package app

import (
	"context"

	"github.com/tn-go-course/lynks/memcache/internal/models"
)

func (app *App) StoreUrl(ctx context.Context, el *models.StorageElement) (*models.StorageElement, error) {
	err := app.redis.SetString(ctx, el.Key, el.Value)
	if err != nil {
		return nil, err
	}
	return el, nil
}

func (app *App) GetUrl(ctx context.Context, stringId string) (string, error) {
	res, err := app.redis.GetString(ctx, stringId)
	if err != nil {
		return "", err
	}
	return res, nil
}
