package app

import "context"

// функция для сохранения в кеше, для запуска в го рутине
func (app *App) cacheString(_ context.Context, key string, value string) {
	defer func() {
		if err := recover(); err != nil {
			app.logger.Error().Caller().Msg("store to cache panic")
		}
	}()

	err := app.cache.SetString(context.Background(), key, value)
	if err != nil {
		app.logger.Error().Err(err).Caller().Msg("store to cache failed")
	}
}
