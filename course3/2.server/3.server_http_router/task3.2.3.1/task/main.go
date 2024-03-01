package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/1", helloHandler1)
	r.Get("/2", helloHandler2)
	r.Get("/3", helloHandler3)

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
