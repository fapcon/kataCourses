package main

import (
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

type DineInOrder struct {
	orderDetails map[string]int
}
type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (d *DineInOrder) AddItem(item string, quantity int) error {
	if quantity > 0 {
		d.orderDetails[item] = quantity
	} else {
		return fmt.Errorf("q <= 0")
	}
	return nil
}
func (t *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity > 0 {
		t.orderDetails[item] = quantity
	} else {
		return fmt.Errorf("q <= 0")
	}
	return nil
}

func (d *DineInOrder) RemoveItem(item string) error {
	_, ok := d.orderDetails[item]
	if ok {
		delete(d.orderDetails, item)
	} else {
		return fmt.Errorf("item not exists")
	}
	return nil
}
func (t *TakeAwayOrder) RemoveItem(item string) error {
	_, ok := t.orderDetails[item]
	if ok {
		delete(t.orderDetails, item)
	} else {
		return fmt.Errorf("item not exists")
	}
	return nil
}

func (d *DineInOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}
func (t *TakeAwayOrder) GetOrderDetails() map[string]int {
	return t.orderDetails
}

func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	ManageOrder(dineIn)
	ManageOrder(takeAway)
}
