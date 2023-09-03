package api

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
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
		reqId := rand.Intn(1_000_000)
		ts := time.Now().Unix()
		ctx := context.WithValue(r.Context(), requestIdContextKey, reqId)
		ctx = context.WithValue(ctx, requestTsKey, ts)
		ctx = a.logger.With().Int("requestId", reqId).Int64("reqTs", ts).Logger().WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Api) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		url := r.URL.Path
		method := r.Method
		zerolog.Ctx(r.Context()).Info().Str("method", method).Str("url", url).Str("ip", ip).Msg("incomig request")
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
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
				a.logger.Error().Caller().Stack().Msg("internal error")
				returnError(ctx, http.StatusInternalServerError, errInternal, w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
