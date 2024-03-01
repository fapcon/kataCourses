package main

import (
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Complete bool
}

var orders []*Order
var completeOrders map[int]bool
var wg sync.WaitGroup
var processTimes chan time.Duration
var sinceProgramStarted time.Duration
var count int
var limitCount int

func main() {
	count = 30
	limitCount = 5
	processTimes = make(chan time.Duration, count)
	orders = GenerateOrders(count)
	completeOrders = GenerateCompleteOrders(count)
	programStart := time.Now()
	LimitSpawnOrderProcessing(limitCount)

	wg.Wait()
	sinceProgramStarted = time.Since(programStart)
	go func() {
		time.Sleep(1 * time.Second)
		close(processTimes)
	}()
	checkTimeDifference(limitCount)
}

func checkTimeDifference(limitCount int) {
	// do not edit
	var averageTime time.Duration
	var orderProcessTotalTime time.Duration
	var orderProcessedCount int
	for v := range processTimes {
		orderProcessedCount++
		orderProcessTotalTime += v
	}
	if orderProcessedCount != count {
		panic("orderProcessedCount != count")
	}
	averageTime = orderProcessTotalTime / time.Duration(orderProcessedCount)
	println("orderProcessTotalTime", orderProcessTotalTime/time.Second)
	println("averageTime", averageTime/time.Second)
	println("sinceProgramStarted", sinceProgramStarted/time.Second)
	println("sinceProgramStarted average", sinceProgramStarted/(time.Duration(orderProcessedCount)*time.Second))
	println("orderProcessTotalTime - sinceProgramStarted", (orderProcessTotalTime-sinceProgramStarted)/time.Second)
	if (orderProcessTotalTime/time.Duration(limitCount)-sinceProgramStarted)/time.Second > 0 {
		panic("(orderProcessTotalTime-sinceProgramStarted)/time.Second > 0")
	}
}

func LimitSpawnOrderProcessing(limitCount int) {
	limit := make(chan struct{}, limitCount)
	var t time.Time
	for i := 0; i < count; i++ {
		limit <- struct{}{}
		wg.Add(1)
		go OrderProcessing(orders[i], limit, t)
	}
	wg.Wait()
}

func OrderProcessing(order *Order, limit chan struct{}, t time.Time) {
	// complete orders if they completed
	<-limit
	time.Sleep(1 * time.Second)
	processTimes <- time.Since(t)
	wg.Done()
}

func GenerateOrders(count int) []*Order {
	// generate uncomplete orders by count variable
	ord := make([]Order, count)
	res := make([]*Order, count)
	for i := 0; i < count; i++ {
		ord[i] = Order{
			ID:       i,
			Complete: false,
		}
	}
	for i := 0; i < count; i++ {
		res[i] = &ord[i]
	}
	return res
}

func GenerateCompleteOrders(maxOrderID int) map[int]bool {
	// chance 50% to generate map of complete order
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < 0.5 {
		// do something with a 50% chance
		for i := 0; i < maxOrderID; i++ {
			if orders[i].Complete == true {
				completeOrders[i] = true
			}
		}
	}
	return completeOrders
}
