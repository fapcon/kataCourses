package geo

import (
	"log"
	"rpc/internal/models"
	"rpc/internal/service"
)

type ServerGeo struct {
	geo service.GeoProvide
}

func (g *ServerGeo) Search(args string, reply *[]byte) error {
	address, err := g.geo.Search(args)
	if err != nil {
		log.Fatal("err call rpc:", err)
	}
	*reply = address
	return nil
}

func (g *ServerGeo) Geocode(inp *models.GeocodeRequest, reply *[]byte) error {
	address, err := g.geo.Geocode(inp)
	if err != nil {
		log.Fatal("err call rpc:", err)
	}
	*reply = address
	return nil
}
