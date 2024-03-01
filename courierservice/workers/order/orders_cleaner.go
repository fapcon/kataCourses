package order

import (
	"context"
	"courierservice/module/order/service"
	"fmt"
	"log"
	"time"
)

const (
	orderCleanInterval = 5 * time.Second
)

// OrderCleaner воркер, который удаляет старые заказы
// используя метод orderService.RemoveOldOrders()
type OrderCleaner struct {
	orderService service.Orderer
}

func NewOrderCleaner(orderService service.Orderer) *OrderCleaner {
	return &OrderCleaner{orderService: orderService}
}

func (o *OrderCleaner) Run() {
	// исользовать горутину и select
	// внутри горутины нужно использовать time.NewTicker()
	// и вызывать метод orderService.RemoveOldOrders()
	// если при удалении заказов произошла ошибка, то нужно вывести ее в лог

	go func() {
		var err error
		t := time.NewTicker(orderCleanInterval)
		for {
			<-t.C
			err = o.orderService.RemoveOldOrders(context.Background())
			if err != nil {
				log.Fatal("order cleaning error:" + err.Error())
				fmt.Println("clean worker - orders cleaning error:" + err.Error())
				continue
			}
			//fmt.Println("clean worker - orders cleaned")
		}

	}()
}
