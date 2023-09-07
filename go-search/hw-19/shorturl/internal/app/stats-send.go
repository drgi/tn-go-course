package app

import (
	"context"
	"encoding/json"
	"log"

	"github.com/tn-go-course/go-search/hw-19/shorturl/pkg/db"
)

func (app *App) SendStat(ctx context.Context, u *db.Url) {
	b, err := json.Marshal(u)
	if err != nil {
		log.Println("marashal message failed: ", err)
	}
	app.broker.Send(ctx, b)
}
