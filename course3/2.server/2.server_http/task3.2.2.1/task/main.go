package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, request *http.Request) {
		rw.Write([]byte("HTTP/1.1 200 OK\r\nHello World!"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
