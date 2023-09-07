package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

// Тут решил сделать функции и структуру обертку для ответов АПИ
type ResponseEnvelope struct {
	Result interface{} `json:"data"`
	Error  *ApiError   `json:"error"`
}

type ApiError struct {
	RequestId int    `json:"requestId"`
	Message   string `json:"message"`
}

func (a *Api) returnData(ctx context.Context, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	response := &ResponseEnvelope{Result: data}
	b, _ := json.Marshal(response)
	w.Write(b)
	a.SendMetrics(ctx)
}

func (a *Api) returnError(ctx context.Context, code int, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	reqId := ctx.Value("reqId").(int)
	zerolog.Ctx(ctx).Error().Err(err).Int("httpCode", code).Msg("handle request failed")
	response := &ResponseEnvelope{
		Error: &ApiError{
			RequestId: reqId,
			Message:   err.Error(),
		}}
	b, _ := json.Marshal(response)
	w.Write(b)
	a.SendMetrics(ctx)
}

func (a *Api) returnRedirect(w http.ResponseWriter, r *http.Request, url string) {
	_ = r.Context()
	http.Redirect(w, r, url, http.StatusSeeOther)
	a.SendMetrics(r.Context())
}
