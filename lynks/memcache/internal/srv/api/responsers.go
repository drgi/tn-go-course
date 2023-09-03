package api

import (
	"context"
	"encoding/json"
	"net/http"
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

func returnData(ctx context.Context, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	response := &ResponseEnvelope{Result: data}
	b, _ := json.Marshal(response)
	w.Write(b)
}

func returnError(ctx context.Context, code int, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	reqId := ctx.Value("reqId").(int)
	response := &ResponseEnvelope{
		Error: &ApiError{
			RequestId: reqId,
			Message:   err.Error(),
		}}
	b, _ := json.Marshal(response)
	w.Write(b)
}

func returnRedirect(w http.ResponseWriter, r *http.Request, url string) {
	_ = r.Context()
	http.Redirect(w, r, url, http.StatusSeeOther)
}
