package order

import (
	"context"
	"courierservice/module/order/service"
	"fmt"
	"log"
	"time"
)

const (
	// order generation interval
	orderGenerationInterval = 10 * time.Millisecond
	maxOrdersCount          = 200
)

// worker generates orders and put them into redis
type OrderGenerator struct {
	orderService service.Orderer
}

func NewOrderGenerator(orderService service.Orderer) *OrderGenerator {
	return &OrderGenerator{orderService: orderService}
}

func (o *OrderGenerator) Run() {
	// запускаем горутину, которая будет генерировать заказы не более чем раз в 10 миллисекунд
	// не более 200 заказов используя константы orderGenerationInterval и maxOrdersCount
	// нужно использовать метод orderService.GetCount() для получения количества заказов
	// и метод orderService.GenerateOrder() для генерации заказа
	// если количество заказов меньше maxOrdersCount, то нужно сгенерировать новый заказ
	// если количество заказов больше или равно maxOrdersCount, то не нужно ничего делать
	// если при генерации заказа произошла ошибка, то нужно вывести ее в лог
	// если при получении количества заказов произошла ошибка, то нужно вывести ее в лог
	// внутри горутины нужно использовать select и time.NewTicker()

	go func() {
		t := time.NewTicker(orderGenerationInterval)
		defer t.Stop()
		ctx := context.Background()

		for {
			<-t.C
			n, err := o.orderService.GetCount(ctx)
			if err != nil {
				log.Println("order counting error:" + err.Error())
				fmt.Println("order generating error:" + err.Error())
				continue

			}

			if n < maxOrdersCount {
				err = o.orderService.GenerateOrder(ctx)
				if err != nil {
					log.Println("order generating error:" + err.Error())
					fmt.Println("order generating error:" + err.Error())
					continue

				}
				//fmt.Println("order generated:", n)
			}

		}
	}()

}
