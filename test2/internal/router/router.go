package router

import (
	"fmt"
	"html/template"
	"net/http"

	"postgrerepository/internal/controller" //"usr/local/go/src/postgrerepository/internal/controller" //
	//"usr/local/go/src/postgrerepository/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//var token *jwtauth.JWTAuth

func NewApiRouter(ctrl *controller.Controller) *chi.Mux {
	r := chi.NewRouter()
	//token = jwtauth.New("HS256", []byte("secret"), nil)
	r.Use(middleware.Logger)
	r.Get("/api/docs*", swaggerUI)
	r.Post("/api/users/create", ctrl.User.Create)
	r.Get("/api/users/{id}", ctrl.User.GetByID)
	r.Post("/api/users/update/{id}", ctrl.User.Update)
	r.Post("/api/users/delete/{id}", ctrl.User.Delete)
	r.Get("/api/users/list", ctrl.User.List)
	r.Get("/docs/*", swaggFileServe)

	return r

}
func swaggerUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templ, err := template.New("swagger").Parse(swaggerTemplate)
	if err != nil {
		fmt.Println("swagger template parsing error:", err.Error())
		return
	}

	err = templ.Execute(w, "")
	if err != nil {
		fmt.Println("template execution error:", err.Error())
		return
	}
}

func swaggFileServe(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "/docs/swagger.json")
}

const (
	swaggerTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-standalone-preset.js"></script>
		<!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-standalone-preset.js"></script> -->
		<script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
		<!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-bundle.js"></script> -->
		<link rel="stylesheet" href="//unpkg.com/swagger-ui-dist@3/swagger-ui.css" />
		<!-- <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui.css" /> -->
		<style>
			body {
				margin: 0;
			}
		</style>
		<title>Swagger</title>
	</head>
	<body>
		<div id="swagger-ui"></div>
		<script>
			window.onload = function() {
			  SwaggerUIBundle({
				url: "/docs/swagger.json?",
				dom_id: '#swagger-ui',
				presets: [
				  SwaggerUIBundle.presets.apis,
				  SwaggerUIStandalonePreset
				],
				layout: "StandaloneLayout"
			  })
			}
		</script>
	</body>
	</html>
	`
)
