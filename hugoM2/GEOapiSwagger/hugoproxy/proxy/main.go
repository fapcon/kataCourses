package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// @title swagger API
// @version 1.0
// @description API geogrpc

// @host localhost:1313
// @BasePath /

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/swagger", SwaggerUI)
	r.Get("/*", func(writer http.ResponseWriter, request *http.Request) {
		http.StripPrefix("/swagger", http.FileServer(http.Dir("./swagger"))).ServeHTTP(writer, request)
	})
	r.Post("/address/search", searchHandler)
	r.Post("/api/address/geocode", geocodeHandler)

	http.ListenAndServe(":8080", r)
}
