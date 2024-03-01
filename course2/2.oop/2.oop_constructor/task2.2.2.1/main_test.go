package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	type args struct {
		id      int
		options []OrderOption
	}
	tests := []struct {
		name string
		args args
		want *Order
	}{
		{
			name: "c1",
			args: args{id: 1, options: []OrderOption{WithCustomerID("123"), WithItems([]string{"item1", "item2"}), WithOrderDate(time.Now())}},
			want: &Order{ID: 1, CustomerID: "123", Items: []string{"item1", "item2"}, OrderDate: time.Now()},
		},
		{
			name: "c2",
			args: args{id: 3, options: []OrderOption{WithCustomerID("123"), WithItems([]string{"item1", "item2"}), WithOrderDate(time.Now())}},
			want: &Order{ID: 1, CustomerID: "123", Items: []string{"item1", "item2"}, OrderDate: time.Now()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrder(tt.args.id, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithCustomerID(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want OrderOption
	}{
		{
			name: "c1",
			args: args{value: "123"},
			want: func(order *Order) { order.CustomerID = "123" },
		},
		{
			name: "c2",
			args: args{value: "123456"},
			want: func(order *Order) { order.CustomerID = "123" },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithCustomerID(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCustomerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithItems(t *testing.T) {
	type args struct {
		value []string
	}
	tests := []struct {
		name string
		args args
		want OrderOption
	}{
		{
			name: "c1",
			args: args{value: []string{"item1", "item2"}},
			want: func(order *Order) { order.Items = []string{"item1", "item2"} },
		},
		{
			name: "c2",
			args: args{value: []string{"item1", "item2"}},
			want: func(order *Order) { order.Items = []string{"item1"} },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithItems(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithOrderDate(t *testing.T) {
	type args struct {
		value time.Time
	}
	tests := []struct {
		name string
		args args
		want OrderOption
	}{
		{
			name: "c1",
			args: args{value: time.Now()},
			want: func(order *Order) { order.OrderDate = time.Now() },
		},
		{
			name: "c2",
			args: args{value: time.Now()},
			want: func(order *Order) { order.OrderDate = time.Now().Add(time.Second * 100) },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithOrderDate(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithOrderDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
