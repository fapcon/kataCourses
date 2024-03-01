package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
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
	}{
		{
			name: "Case1",
			fields: fields{
				client: &http.Client{},
				url:    "https://api.exmo.com/v1",
			},
			args: args{
				pair:  "BTC_USD",
				limit: 30,
				start: time.Now().Add(-time.Hour * 24),
				end:   time.Now(),
			},
			want:    CandlesHistory{},
			wantErr: false,
		},
		{
			name: "Case2",
			fields: fields{
				client: &http.Client{},
				url:    "https://api.exmo.com/v1",
			},
			args: args{
				pair:  "BTC_USD",
				limit: 30,
				start: time.Now().Add(-time.Hour * 24),
				end:   time.Now(),
			},
			want:    CandlesHistory{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := NewMockExchanger(t)
			if tt.wantErr {
				ex.On("GetCandlesHistory", tt.args.pair, tt.args.limit, tt.args.start, tt.args.end).
					Return(CandlesHistory{}, errors.New("errEX"))
			} else {
				ex.On("GetCandlesHistory", tt.args.pair, tt.args.limit, tt.args.start, tt.args.end).
					Return(tt.want, nil)
			}
			got, err := ex.GetCandlesHistory(tt.args.pair, tt.args.limit, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCandlesHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)

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
	}{
		{
			name: "Case1",
			fields: fields{
				client: &http.Client{},
				url:    "https://api.exmo.com/v1",
			},
			args: args{
				pair:  "BTC_USD",
				limit: 5,
				start: time.Now().Add(-2),
				end:   time.Now(),
			},
			want:    make([]float64, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := NewMockExchanger(t)
			ex.On("GetClosePrice", tt.args.pair, tt.args.limit, tt.args.start, tt.args.end).
				Return(tt.want, nil)
			got, err := ex.GetClosePrice(tt.args.pair, tt.args.limit, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClosePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
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
	}{
		{
			name: "Case1",
			fields: fields{
				client: &http.Client{},
				url:    "https://api.exmo.com/v1",
			},
			want:    Currencies{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewMockExchanger(t)
			e.On("GetCurrencies").Return(Currencies{}, nil)
			got, err := e.GetCurrencies()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
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
	}{
		{
			name: "Case1",
			fields: fields{
				client: &http.Client{},
				url:    "https://api.exmo.com/v1",
			},
			args: args{
				limit: 5,
				pairs: []string{"BTC_USD"},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := NewMockExchanger(t)
			ex.On("GetOrderBook", 5, "BTC_USD").Return(OrderBook{}, nil)
			_, err := ex.GetOrderBook(tt.args.limit, tt.args.pairs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("errt")
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
	}{
		{
			name: "Case1",
			fields: fields{
				client: &http.Client{},
				url:    "https://api.exmo.com/v1",
			},
			args: args{
				pairs: []string{"BTC_USD"},
			},
			want:    Trades{},
			wantErr: false,
		},
		{
			name: "Case2",
			fields: fields{
				client: &http.Client{},
				url:    "https://api.exmo.com/v1",
			},
			args: args{
				pairs: []string{"BTC_USD"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := NewMockExchanger(t)
			if !tt.wantErr {
				ex.On("GetTrades", "BTC_USD").Return(Trades{}, nil)
				got, err := ex.GetTrades(tt.args.pairs...)
				if (err != nil) != tt.wantErr {
					t.Errorf("got:%v want:%v", err, tt.wantErr)
				}
				assert.Equal(t, tt.want, got)
			}
			if tt.wantErr {
				ex.On("GetTrades", "BTC_USD").Return(nil, errors.New("errors"))
				got, err := ex.GetTrades(tt.args.pairs...)
				if (err != nil) != tt.wantErr {
					t.Errorf("got:%v want:%v", err, tt.wantErr)
				}
				assert.Equal(t, tt.want, got)
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
	}{
		{
			name: "Case1",
			fields: fields{
				client: &http.Client{},
				url:    "https://api.exmo.com/v1",
			},
			want:    Ticker{},
			wantErr: false,
		},
		{
			name: "Case2",
			fields: fields{
				client: &http.Client{},
				url:    "https://api.exmo.com/v1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := NewMockExchanger(t)
			if !tt.wantErr {
				ex.On("GetTicker").Return(tt.want, nil)
				got, err := ex.GetTicker()
				if (err != nil) != tt.wantErr {
					t.Errorf("err")
				}
				assert.Equal(t, tt.want, got)
			}
			if tt.wantErr {
				ex.On("GetTicker").Return(tt.want, errors.New("teer"))
				got, err := ex.GetTicker()
				if (err != nil) != tt.wantErr {
					t.Errorf("err")
				}
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestIndicator_EMA(t *testing.T) {

	type args struct {
		period int
		data   []float64
		pair   string
		limit  int
		from   time.Time
		to     time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "Case1",
			args: args{
				period: 5,
				data:   []float64{100, 200, 300, 400, 500},
				pair:   "BTC_USD",
				limit:  30,
				from:   time.Now().Add(-2),
				to:     time.Now(),
			},
			want:    []float64{100, 133, 189, 259, 339},
			wantErr: false,
		},
		{
			name: "Case2",
			args: args{
				period: 5,
				data:   []float64{100, 200, 300, 400, 500},
				pair:   "BTC_USD",
				limit:  30,
				from:   time.Now().Add(-2),
				to:     time.Now(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exchange := NewMockExchanger(t)
			if tt.wantErr {
				exchange.On("GetClosePrice", tt.args.pair, tt.args.limit, tt.args.from, tt.args.to).
					Return(nil, errors.New("errEX"))
				i := NewIndicator(exchange, WithEMA(calculateEMA))
				got, err := i.EMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
				if (err != nil) != tt.wantErr {
					t.Errorf("err")
				}
				assert.Equal(t, tt.want, got)
			}
			if !tt.wantErr {
				exchange.On("GetClosePrice", tt.args.pair, tt.args.limit, tt.args.from, tt.args.to).
					Return(tt.args.data, nil)
				i := NewIndicator(exchange, WithEMA(calculateEMA))
				got, _ := i.EMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestIndicator_SMA(t *testing.T) {

	type args struct {
		pair   string
		limit  int
		period int
		from   time.Time
		to     time.Time
		data   []float64
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "Case1",
			args: args{
				pair:   "BTC_USD",
				limit:  30,
				period: 5,
				from:   time.Now().Add(-2),
				to:     time.Now(),
				data:   []float64{100, 200, 300, 400, 500},
			},
			want:    []float64{300},
			wantErr: false,
		},
		{
			name: "Case2",
			args: args{
				pair:   "BTC_USD",
				limit:  30,
				period: 5,
				from:   time.Now().Add(-2),
				to:     time.Now(),
				data:   []float64{100, 200, 300, 400, 500},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exchan := NewMockExchanger(t)
			if !tt.wantErr {
				exchan.On("GetClosePrice", tt.args.pair, tt.args.limit, tt.args.from, tt.args.to).
					Return(tt.args.data, nil)
				i := NewIndicator(exchan, WithSMA(calculateSMA))
				got, err := i.SMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
				if (err != nil) != tt.wantErr {
					t.Errorf("err")
				}
				assert.Equal(t, tt.want, got)
			}
			if tt.wantErr {
				exchan.On("GetClosePrice", tt.args.pair, tt.args.limit, tt.args.from, tt.args.to).
					Return(nil, errors.New("errEX"))
				i := NewIndicator(exchan, WithSMA(calculateSMA))
				got, err := i.SMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
				if (err != nil) != tt.wantErr {
					t.Errorf("err")
				}
				assert.Equal(t, tt.want, got)
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
	}{
		{
			name: "Case1",
			args: args{opts: []func(exmo *Exmo){WithClient(&http.Client{})}},
			want: &Exmo{
				client: &http.Client{},
			},
		},
		{
			name: "Case2",
			args: args{opts: []func(exmo *Exmo){WithURL("https://api.exmo.com/v1")}},
			want: &Exmo{url: "https://api.exmo.com/v1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExmo(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExmo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIndicator(t *testing.T) {
	type args struct {
		exchange Exchanger
		opts     []IndicatorOption
	}

	tests := []struct {
		name string
		args args
		want *Indicator
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

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
	}{
		{
			name: "Case1",
			args: args{client: &http.Client{}},
			want: func(exmo *Exmo) {
				exmo.client = &http.Client{}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantExmo := &Exmo{}
			tt.want(wantExmo)

			gotExmo := &Exmo{}
			WithClient(tt.args.client)(gotExmo)

			if !reflect.DeepEqual(gotExmo, wantExmo) {
				t.Errorf("WithClient() = %v, want %v", gotExmo, wantExmo)
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
	}{
		{
			name: "Case1",
			args: args{url: "https://api.exmo.com/v1"},
			want: func(exmo *Exmo) {
				exmo.url = "https://api.exmo.com/v1"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := &Exmo{}
			tt.want(want)
			got := &Exmo{}
			WithURL(tt.args.url)(got)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("WithClient() = %v, want %v", got, want)
			}
		})
	}
}

func Test_calculateEMA(t *testing.T) {
	type args struct {
		data   []float64
		period int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Case1",
			args: args{
				data:   []float64{100, 200, 300, 400, 500},
				period: 5,
			},
			want: []float64{100, 133, 189, 259, 339},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateEMA(tt.args.data, tt.args.period); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateEMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSMA(t *testing.T) {
	type args struct {
		data   []float64
		period int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Case1",
			args: args{
				data:   []float64{100, 200, 300, 400, 500},
				period: 5,
			},
			want: []float64{300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSMA(tt.args.data, tt.args.period); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateSMA() = %v, want %v", got, tt.want)
			}
		})
	}
}
