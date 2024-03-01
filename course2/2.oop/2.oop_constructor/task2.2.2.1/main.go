package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func WithCustomerID(value string) OrderOption {
	return func(order *Order) {
		order.CustomerID = value
	}
}

func WithItems(value []string) OrderOption {
	return func(order *Order) {
		order.Items = value
	}
}

func WithOrderDate(value time.Time) OrderOption {
	return func(order *Order) {
		order.OrderDate = value
	}
}

func NewOrder(id int, options ...OrderOption) *Order {
	ord := &Order{ID: id}
	for _, option := range options {
		option(ord)
	}
	return ord
}

func main() {
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))

	fmt.Printf("Order: %+v\n", order)
}
