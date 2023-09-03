package api

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

const (
	requestIdContextKey = "reqId"
	requestTsKey        = "requestTs"
	requestTimeout      = 20
)

var (
	errInternal = errors.New("iternal error")
)

func (a *Api) setRequestIdAndTs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), requestIdContextKey, rand.Intn(1_000_000))
		ctx = context.WithValue(ctx, requestTsKey, time.Now().Unix())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Api) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId := r.Context().Value(requestIdContextKey).(int)
		ip := r.RemoteAddr
		url := r.URL.Path
		method := r.Method
		a.logger.Info().Int("requestId", reqId).Str("method", method).Str("url", url).Str("ip", ip).Msg("incomig request")
		next.ServeHTTP(w, r)
	})
}

func (a *Api) setTimeoutAndRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, time.Second*requestTimeout)
		defer func() {
			cancel()
			if r := recover(); r != nil {
				a.logger.Error().Stack().Msg("internal error")
				returnError(ctx, http.StatusInternalServerError, errInternal, w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
