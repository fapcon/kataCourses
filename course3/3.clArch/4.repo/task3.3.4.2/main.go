package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"petstore/internal/controller"
	"petstore/internal/repository"
	"petstore/internal/router"
	"time"

	"github.com/joho/godotenv"
)

// @title API Title
// @version 1.0
// @description This is a sample geogrpc.

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env loading error")
	}

	db := repository.NewPostgreDB()

	defer db.Close()

	ctrl := controller.NewController(db)

	r := router.NewRouter(ctrl)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Starting geogrpc...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
