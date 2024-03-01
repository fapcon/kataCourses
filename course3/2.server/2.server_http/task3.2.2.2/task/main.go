package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	e := godotenv.Load("./.env") //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, request *http.Request) {
		rw.Write([]byte("HTTP/1.1 200 OK\r\nHello World!"))
	})
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, nil))
}
