package main

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestExmo_GetCandlesHistory(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		pair  string
		limit int
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    CandlesHistory
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Exmo{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := e.GetCandlesHistory(tt.args.pair, tt.args.limit, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCandlesHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCandlesHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_GetClosePrice(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		pair  string
		limit int
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []float64
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Exmo{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := e.GetClosePrice(tt.args.pair, tt.args.limit, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClosePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClosePrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_GetCurrencies(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    Currencies
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Exmo{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := e.GetCurrencies()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrencies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_GetOrderBook(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		limit int
		pairs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    OrderBook
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Exmo{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := e.GetOrderBook(tt.args.limit, tt.args.pairs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrderBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrderBook() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_GetTicker(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    Ticker
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Exmo{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := e.GetTicker()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_GetTrades(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		pairs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Trades
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Exmo{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := e.GetTrades(tt.args.pairs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTrades() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTrades() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_doPostRequest(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		urll string
		data url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Exmo{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := e.doPostRequest(tt.args.urll, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("doPostRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doPostRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExmo_doRequest(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Exmo{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := e.doRequest(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewExmo(t *testing.T) {
	type args struct {
		opts []func(exmo *Exmo)
	}
	tests := []struct {
		name string
		args args
		want *Exmo
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExmo(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExmo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithClient(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name string
		args args
		want func(exmo *Exmo)
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want func(exmo *Exmo)
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithURL(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
