package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	dadata "github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

func main() {
	//
	//creds := client.Credentials{
	//	ApiKeyValue:    "c4774295a9b0dc55c9b8e73cb299d5c5faeb2da8",
	//	SecretKeyValue: "efefac572c92aaddf6fe54d04a31f00d72702ed1",
	//}
	//
	//api := dadata.NewCleanApi(client.WithCredentialProvider(&creds))
	//query := "королев пионерская 4"
	//result, err := api.Address(context.Background(), query)
	//if err != nil {
	//	fmt.Errorf("zapros minus")
	//}
	//
	//for _, addr := range result {
	//	fullAddr := &Address{GeoLon: addr.GeoLon, GeoLat: addr.GeoLat, Source: addr.Source, Result: addr.Result, PostalCode: addr.PostalCode, Country: addr.Country, Region: addr.Region, CityArea: addr.CityArea, CityDistrict: addr.CityDistrict, Street: addr.Street, House: addr.House}
	//	fmt.Println(fullAddr)
	//}

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1313/swagger/doc.json"), //The url pointing to API definition
	))
	//r.Get("/swagger/index.html", SwaggerUI)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/address/search", searchHandler)

	r.Post("/api/address/geocode", geocodeHandler)

	http.ListenAndServe(":1313", r)
}

const (
	DadataURL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs"
	APIKey    = "c4774295a9b0dc55c9b8e73cb299d5c5faeb2da8"
)

//type Address struct {
//	Value             string                 `json:"value"`
//	UnrestrictedValue string                 `json:"unrestricted_value"`
//	Data              map[string]interface{} `json:"data"`
//}

type Address struct {
	Source       string `json:"source"`
	Result       string `json:"result"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	CityArea     string `json:"city_area"`
	CityDistrict string `json:"city_district"`
	Street       string `json:"street"`
	House        string `json:"house"`
	GeoLat       string `json:"geo_lat"`
	GeoLon       string `json:"geo_lon"`
	QcGeo        int64  `json:"qc_geo"`
}

type SearchRequest struct {
	Query string `json:"query"`
}

type SearchResponse struct {
	Addresses []*Address `json:"suggestions"`
}

type GeocodeRequest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lon"`
}

type GeocodeResponse struct {
	Addresses []*Address `json:"suggestions"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var fullAddr *Address
	var req SearchRequest

	creds := client.Credentials{
		ApiKeyValue:    "c4774295a9b0dc55c9b8e73cb299d5c5faeb2da8",
		SecretKeyValue: "efefac572c92aaddf6fe54d04a31f00d72702ed1",
	}

	api := dadata.NewCleanApi(client.WithCredentialProvider(&creds))
	//query := "королев пионерская 4"

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	result, err := api.Address(context.Background(), req.Query)
	if err != nil {
		fmt.Errorf("zapros minus")
	}

	for _, addr := range result {
		fullAddr = &Address{GeoLon: addr.GeoLon, GeoLat: addr.GeoLat, Source: addr.Source, Result: addr.Result, PostalCode: addr.PostalCode, Country: addr.Country, Region: addr.Region, CityArea: addr.CityArea, CityDistrict: addr.CityDistrict, Street: addr.Street, House: addr.House}
		fmt.Println(fullAddr)
	}
	//var req SearchRequest
	//
	//err = json.NewDecoder(r.Body).Decode(&req)
	//if err != nil {
	//	http.Error(w, "Invalid request format", http.StatusBadRequest)
	//	return
	//}
	//
	url := DadataURL + "/suggest/address"
	//reqMap := map[string]interface{}{"query": req.Query}
	body, err := json.Marshal(fullAddr)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	dadataReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	dadataReq.Header.Set("Content-Type", "application/json")
	dadataReq.Header.Set("Accept", "application/json")
	dadataReq.Header.Set("Authorization", "Token "+APIKey)

	dadataResp, err := client.Do(dadataReq)
	if err != nil {
		http.Error(w, "Failed to connect to Dadata service", http.StatusInternalServerError)
		return
	}
	defer dadataResp.Body.Close()

	if dadataResp.StatusCode != http.StatusOK {
		http.Error(w, "Dadata service error", dadataResp.StatusCode)
		return
	}

	var resp SearchResponse
	err = json.NewDecoder(dadataResp.Body).Decode(&resp)
	if err != nil {
		http.Error(w, "Failed to parse Dadata response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Addresses)
}

func geocodeHandler(w http.ResponseWriter, r *http.Request) {
	var req GeocodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	url := DadataURL + "/geolocate/address"
	reqMap := map[string]interface{}{"lat": req.Lat, "lon": req.Lng}
	body, err := json.Marshal(reqMap)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	dadataReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	dadataReq.Header.Set("Content-Type", "application/json")
	dadataReq.Header.Set("Accept", "application/json")
	dadataReq.Header.Set("Authorization", "Token "+APIKey)

	dadataResp, err := client.Do(dadataReq)
	if err != nil {
		http.Error(w, "Failed to connect to Dadata service", http.StatusInternalServerError)
		return
	}
	defer dadataResp.Body.Close()

	if dadataResp.StatusCode != http.StatusOK {
		http.Error(w, "Dadata service error", dadataResp.StatusCode)
		return
	}

	var resp GeocodeResponse
	err = json.NewDecoder(dadataResp.Body).Decode(&resp)
	if err != nil {
		http.Error(w, "Failed to parse Dadata response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Addresses)
}
