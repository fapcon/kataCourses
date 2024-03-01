package router

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
	"petstore/internal/controller"
)

func NewRouter(ctrl *controller.Controller) http.Handler {
	r := chi.NewRouter()

	r.Use(Logger)
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Route("/api", func(r chi.Router) {
		//r.Get("/swagger/*", httpSwagger.WrapHandler)
		r.Post("/pet", ctrl.PetCtrl.Create)
		r.Put("/pet", ctrl.PetCtrl.Update)
		r.Get("/pet/findByStatus", ctrl.PetCtrl.GetByStatus)
		r.Get("/pet/{petId}", ctrl.PetCtrl.GetById)
		r.Put("/pet/{petId}", ctrl.PetCtrl.UpdateById)
		r.Delete("/pet/{petId}", ctrl.PetCtrl.DeleteById)

		r.Route("/store", func(r chi.Router) {
			r.Post("/order", ctrl.StCtrl.Create)
			r.Get("/order/{OrderId}", ctrl.StCtrl.GetByID)
			r.Delete("/order/{OrderId}", ctrl.StCtrl.DeleteByID)
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/{username}", ctrl.UsCtrl.GetByUsername)
			r.Put("/{username}", ctrl.UsCtrl.UpdateByUsername)
			r.Delete("/{username}", ctrl.UsCtrl.DeleteByUsername)
			r.Post("/login", ctrl.UsCtrl.Login)
			r.Post("/logout", ctrl.UsCtrl.Logout)
			r.Post("/createWithArray", ctrl.UsCtrl.CreateWithArray)
			r.Post("/", ctrl.UsCtrl.CreateUser)
		})
	})

	return r

}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l, _ := zap.NewProduction()
		defer l.Sync()
		l.Info("request got",
			zap.String("path", r.URL.Path),
			zap.String("method", r.Method),
			zap.String("IP", r.RemoteAddr))
	})
}
