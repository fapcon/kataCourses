package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/group1/1", helloHandler1g1)
	r.Get("/group1/2", helloHandler1g2)
	r.Get("/group1/3", helloHandler1g3)
	r.Get("/group2/1", helloHandler2g1)
	r.Get("/group2/2", helloHandler2g2)
	r.Get("/group2/3", helloHandler2g3)
	r.Get("/group3/1", helloHandler3g1)
	r.Get("/group3/2", helloHandler3g2)
	r.Get("/group3/3", helloHandler3g3)

	http.ListenAndServe(":8080", r)
}

func helloHandler1g1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nGroup 1 Hello World!"))
}

func helloHandler1g2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nGroup 1 Hello World! 2"))
}
func helloHandler1g3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nGroup 1 Hello World! 3"))
}
func helloHandler2g1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nGroup 2 Hello World!"))
}

func helloHandler2g2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nGroup 2 Hello World! 2"))
}
func helloHandler2g3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nGroup 2 Hello World! 3"))
}
func helloHandler3g1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nGroup 3 Hello World!"))
}

func helloHandler3g2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nGroup 3 Hello World! 2"))
}
func helloHandler3g3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.1 200 OK\nGroup 3 Hello World! 3"))
}
