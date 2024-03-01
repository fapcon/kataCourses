package main

import (
	"google.golang.org/grpc"
	"grpc/internal/grpc/geo"
	grpcpr "grpc/protos/gen"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	server := grpc.NewServer()
	grpcpr.RegisterGeoServiceServer(server, &geo.ServiceGeo{})

	log.Println("Запуск gRPC сервера...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
