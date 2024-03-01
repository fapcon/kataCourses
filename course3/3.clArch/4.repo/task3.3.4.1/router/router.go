package router

import (
	"awesomeProject/internal/handler"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(userHandler *handler.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Post("/api/users", userHandler.Create)
	r.Get("/api/users/{id}", userHandler.GetById)
	r.Put("/api/users/{id}", userHandler.Update)
	r.Delete("/api/users/{id}", userHandler.Delete)

	r.Get("/api/users", userHandler.List)

	return r
}
