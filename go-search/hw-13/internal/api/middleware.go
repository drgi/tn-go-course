package api

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"runtime/debug"
	"time"
)

const (
	reqIdContextKey = "reqId"
	requestTimeout  = 20
)

var (
	errInternal = errors.New("iternal error")
)

func (a *Api) setRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), reqIdContextKey, rand.Intn(1_000_000))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Api) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId := r.Context().Value(reqIdContextKey).(int)
		ip := r.RemoteAddr
		url := r.URL.Path
		method := r.Method
		a.logger.Info("incoming request", "requestId", reqId, "method", method, "url", url, "ip", ip)
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
				a.logger.Error("panic", "stack", string(debug.Stack()))
				returnError(ctx, http.StatusInternalServerError, errInternal, w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
