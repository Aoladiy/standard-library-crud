package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
	"github.com/google/uuid"
)

type CtxKey string

const RequestId CtxKey = "request_id"

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cw := &CustomResponseWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		requestId, ok := r.Context().Value(RequestId).(uuid.UUID)
		if !ok {
			log.Println("Something went wrong - request id is failed converting into uuid")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		start := time.Now()
		next.ServeHTTP(cw, r)
		finish := time.Since(start)
		log.Println(requestId, cw.StatusCode, r.Method, r.URL.Path, finish)
	})
}

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RequestId, uuid.New())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.TimeoutHandler(next, time.Duration(1)*time.Second, "Timeout from middleware expired")
}

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		if envLoading.EnvVars.LoadedUsername == username && envLoading.EnvVars.LoadedPassword == password {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	})
}

func ChainOfMiddleware(handler http.Handler, middleware ...func(next http.Handler) http.Handler) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}
