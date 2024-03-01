package main

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(LoggerMiddleware)
		r.Get("/1", helloHandler1)
	})

	r.Group(func(r chi.Router) {
		r.Use(LoggerMiddleware)
		r.Get("/2", helloHandler2)
	})

	r.Group(func(r chi.Router) {
		r.Use(LoggerMiddleware)
		r.Get("/3", helloHandler3)
	})
	//r.Use(middleware.Logger)
	//r.Get("/1", helloHandler1)
	//r.Get("/2", helloHandler2)
	//r.Get("/3", helloHandler3)

	http.ListenAndServe(":8080", r)
}

func helloHandler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nHello World!"))
}

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nHello World! 2"))
}
func helloHandler3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nHello World! 3"))
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, _ := zap.NewProduction()
		defer logger.Sync() // flushes buffer, if any
		sugar := logger.Sugar()
		sugar.Infow("failed to fetch URL",
			// Structured context as loosely typed key-value pairs.
			"url", r.URL,
			"attempt", 3,
			"backoff", time.Second,
		)
		sugar.Infof("Failed to fetch URL: %s", r.URL)
	})
}
