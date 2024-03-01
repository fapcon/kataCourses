package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"postgrerepository/internal/controller"
	"postgrerepository/internal/repository"
	"postgrerepository/internal/router"

	"github.com/joho/godotenv"
)

// swagger:route POST /api/address/search  addr RequestAddressSearch
// getting address
// responses:
// 200:

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env loading error")
	}

	db := repository.NewPostgreDB()

	defer db.Close()

	repo := repository.NewPostgreUserRepo(db)
	ctrl := controller.NewController(repo)

	r := router.NewApiRouter(ctrl)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//go TimeUpdate()
	//go BinTreeBuilt()
	//go graphRandomBuilt()

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

	//http.ListenAndServe(":8080", router)

}
