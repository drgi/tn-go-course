package api

import (
	"context"
	"time"
)

func (a *Api) SendMetrics(ctx context.Context) {
	url := ctx.Value(requestUrlKey).(string)
	ts := ctx.Value(requestTsKey).(int64)
	a.metric.IncrementRequest(url, time.Unix(ts, 0))
}
