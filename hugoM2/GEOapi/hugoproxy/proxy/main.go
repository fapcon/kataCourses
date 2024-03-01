package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

//

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/address/search", searchHandler)
	r.Post("/api/address/geocode", geocodeHandler)

	http.ListenAndServe(":8080", r)
}
