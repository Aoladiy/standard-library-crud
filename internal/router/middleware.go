package router

import (
	"log"
	"net/http"
	"time"
)

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
		start := time.Now()
		next.ServeHTTP(cw, r)
		finish := time.Since(start)
		log.Println(cw.StatusCode, r.Method, r.URL.Path, finish)
	})
}

func SecondLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cw := &CustomResponseWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		start := time.Now()
		next.ServeHTTP(cw, r)
		finish := time.Since(start)
		log.Println(cw.StatusCode, r.Method, r.URL.Path, finish, "PAY ATTENTION!!! This is second middleware!!!")
	})
}

func RegisterMiddleware(handler http.Handler, middleware ...func(next http.Handler) http.Handler) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}
