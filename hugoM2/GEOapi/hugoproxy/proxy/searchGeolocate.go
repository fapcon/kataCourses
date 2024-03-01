package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	DadataURL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs"
	APIKey    = "c4774295a9b0dc55c9b8e73cb299d5c5faeb2da8"
)

type Address struct {
	Value             string                 `json:"value"`
	UnrestrictedValue string                 `json:"unrestricted_value"`
	Data              map[string]interface{} `json:"data"`
}

// SearchRequest is a model for search request parameters
// swagger:parameters searchHandler
type SearchRequest struct {
	// Query is the address to search for
	// in: body
	// required: true
	Query string `json:"query"`
}

// SearchResponse is a model for search response data
// swagger:response searchResponse
type SearchResponse struct {
	// in: addresses
	Addresses []*Address `json:"suggestions"`
}

// GeocodeRequest is a model for search request parameters
// swagger:parameters geocodeHandler
type GeocodeRequest struct {
	// in: body
	// required: true
	Lat float64 `json:"lat"`
	Lng float64 `json:"lon"`
}

// GeocodeResponse is a model for search response data
// swagger:response geocodeResponse
type GeocodeResponse struct {
	// in: addresses
	Addresses []*Address `json:"suggestions"`
}

// searchHandler retrieves information about an address
// swagger:route POST /api/address/search searchHandler
//
// # Search for an address
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http, https
//
// Responses:
//
//	200: searchResponse
func searchHandler(w http.ResponseWriter, r *http.Request) {
	var req SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	url := DadataURL + "/suggest/address"
	reqMap := map[string]interface{}{"query": req.Query}
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
