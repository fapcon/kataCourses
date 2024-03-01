package main

import (
	"reflect"
	"testing"
)

func TestStatisticProfit_Average(t *testing.T) {
	type fields struct {
		product *Product
	}
	type args struct {
		prices []float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "c1",
			fields: fields{product: &Product{ProductID: ProductCocaCola,
				Sells:        []float64{2.0, 4.0, 5.5},
				Buys:         []float64{1.5, 1.9, 4.6},
				CurrentPrice: 2.66}},
			args: args{prices: []float64{1.0, 2.0, 3.0}},
			want: 2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatisticProfit{
				product: tt.fields.product,
			}
			if got := s.Average(tt.args.prices); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetAllData(t *testing.T) {
	type fields struct {
		product *Product
	}
	tests := []struct {
		name   string
		fields fields
		want   []float64
	}{
		{
			name: "c1",
			fields: fields{product: &Product{ProductID: ProductCocaCola,
				Sells:        []float64{2.0, 4.0, 3.0},
				Buys:         []float64{1.0, 1.0, 1.0},
				CurrentPrice: 2}},
			want: []float64{2, 6, 100, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatisticProfit{
				product: tt.fields.product,
			}
			if got := s.GetAllData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetAverageProfit(t *testing.T) {
	type fields struct {
		product *Product
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "c1",
			fields: fields{product: &Product{ProductID: ProductCocaCola,
				Sells:        []float64{2.0, 4.0, 3.0},
				Buys:         []float64{1.0, 1.0, 1.0},
				CurrentPrice: 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatisticProfit{
				product: tt.fields.product,
			}
			if got := s.GetAverageProfit(); got != tt.want {
				t.Errorf("GetAverageProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetAverageProfitPercent(t *testing.T) {
	type fields struct {
		product *Product
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "c1",
			fields: fields{product: &Product{ProductID: ProductCocaCola,
				Sells:        []float64{2.0, 4.0, 3.0},
				Buys:         []float64{1.0, 1.0, 1.0},
				CurrentPrice: 2}},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatisticProfit{
				product: tt.fields.product,
			}
			if got := s.GetAverageProfitPercent(); got != tt.want {
				t.Errorf("GetAverageProfitPercent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetCurrentProfit(t *testing.T) {
	type fields struct {
		product *Product
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "c1",
			fields: fields{product: &Product{ProductID: ProductCocaCola,
				Sells:        []float64{2.0, 4.0, 3.0},
				Buys:         []float64{1.0, 1.0, 1.0},
				CurrentPrice: 2}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatisticProfit{
				product: tt.fields.product,
			}
			if got := s.GetCurrentProfit(); got != tt.want {
				t.Errorf("GetCurrentProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetDifferenceProfit(t *testing.T) {
	type fields struct {
		product *Product
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "c1",
			fields: fields{product: &Product{ProductID: ProductCocaCola,
				Sells:        []float64{2.0, 4.0, 3.0},
				Buys:         []float64{1.0, 1.0, 1.0},
				CurrentPrice: 2}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatisticProfit{
				product: tt.fields.product,
			}
			if got := s.GetDifferenceProfit(); got != tt.want {
				t.Errorf("GetDifferenceProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_SetProduct(t *testing.T) {
	type fields struct {
		product *Product
	}
	type args struct {
		p *Product
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "c1",
			fields: fields{product: &Product{ProductID: ProductCocaCola,
				Sells:        []float64{2.0, 4.0, 3.0},
				Buys:         []float64{1.0, 1.0, 1.0},
				CurrentPrice: 2}},
			args: args{p: &Product{ProductID: ProductCocaCola,
				Sells:        []float64{2.0, 4.0, 3.0},
				Buys:         []float64{1.0, 1.0, 1.0},
				CurrentPrice: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatisticProfit{
				product: tt.fields.product,
			}
			s.SetProduct(tt.args.p)
		})
	}
}

func TestStatisticProfit_Sum(t *testing.T) {
	type fields struct {
		product *Product
	}
	type args struct {
		prices []float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "c1",
			fields: fields{product: &Product{ProductID: ProductCocaCola,
				Sells:        []float64{2.0, 4.0, 3.0},
				Buys:         []float64{1.0, 1.0, 1.0},
				CurrentPrice: 2}},
			args: args{prices: []float64{1.0, 2.0, 3.0}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatisticProfit{
				product: tt.fields.product,
			}
			if got := s.Sum(tt.args.prices); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
