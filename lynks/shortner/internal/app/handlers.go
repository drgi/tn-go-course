package app

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/url"

	"github.com/tn-go-course/lynks/shortner/internal/models"
)

func (app *App) CreateShortLink(ctx context.Context, destinationUrl string) (*models.Url, error) {
	u, err := url.Parse(destinationUrl)
	if err != nil {
		app.logger.Error().Err(err).Caller().Str("destinationUrl", destinationUrl).Msg("parse url failed")
		return nil, err
	}
	stringId, err := generateStringId(4)
	if err != nil {
		app.logger.Error().Err(err).Caller().Str("destinationUrl", destinationUrl).Msg("generate id failed")
		return nil, err
	}
	newUrl := &models.Url{
		StringID: stringId,
		Url:      u.String(),
	}

	newUrl.ID, err = app.repo.CreateUrl(ctx, newUrl)
	if err != nil {
		app.logger.Error().Err(err).Caller().Msg("create url in DB failed")
		return nil, err
	}

	go app.cacheString(ctx, newUrl.StringID, newUrl.Url)

	return newUrl, nil
}

func (app *App) RestoreLink(ctx context.Context, stringId string) (*models.Url, error) {
	var err error
	url := &models.Url{
		StringID: stringId,
	}
	url.Url, err = app.cache.GetString(ctx, stringId)
	if err != nil {
		app.logger.Error().Err(err).Caller().Str("stringId", stringId).Msg("get from cache failed")
	} else {
		return url, nil
	}
	url, err = app.repo.GetUrlByStringId(ctx, stringId)
	if err != nil {
		app.logger.Error().Err(err).Caller().Str("stringId", stringId).Msg("get url from DB failed")
		return nil, err
	}
	return url, nil
}

func generateStringId(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
