package service

import (
	"fmt"
	"log"
	"os"
	"proxy/internal/models"
	"proxy/internal/rpc/rpcClient"
)

type GeoServicer interface {
	GeoSearch(input string) ([]*models.Address, error)
	GeoCode(lat, lng string) ([]*models.Address, error)
}

type GeoService struct {
}

func NewGeoService() GeoService {
	return GeoService{}
}

func (g *GeoService) GeoSearch(input string) ([]*models.Address, error) {

	protocol := os.Getenv("RPC_PROTOCOL")

	switch protocol {
	case "rpctask":
		rpcFactory := rpcClient.NewClientRpcFactory()
		address, err := rpcFactory.CreateClientAndCall(input)
		if err != nil {
			log.Fatal("err:", err)
			return nil, err
		}
		return address, nil

	case "json-rpctask":
		jsonrpcFactory := rpcClient.NewJsonRpcClientFactory()
		address, err := jsonrpcFactory.CreateClientAndCall(input)
		if err != nil {
			log.Fatal("err:", err)
			return nil, err
		}
		return address, nil
	case "grpc0":
		grpcFactory := rpcClient.NewGrpcClientFactory()
		address, err := grpcFactory.CreateClientAndCall(input)
		if err != nil {
			log.Fatal("err:", err)
		}
		return address, nil

	}
	return nil, fmt.Errorf("unknown rpctask protocol: %s", protocol)
}

func (g *GeoService) GeoCode(lat, lng string) ([]*models.Address, error) {

	return nil, nil
}