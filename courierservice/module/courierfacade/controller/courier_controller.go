package controller

import (
	"context"
	"courierservice/module/courierfacade/service"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type CourierController struct {
	courierService service.CourierFacer
}

func NewCourierController(courierService service.CourierFacer) *CourierController {
	return &CourierController{courierService: courierService}
}

func (c *CourierController) GetStatus(ctx *gin.Context) {
	// установить задержку в 50 миллисекунд
	<-time.NewTimer(50 * time.Millisecond).C

	// получить статус курьера из сервиса courierService используя метод GetStatus
	status := c.courierService.GetStatus(context.Background())
	// отправить статус курьера в ответ
	//fmt.Println("courier status getting:", status)
	ctx.JSON(200, status)
}

func (c *CourierController) MoveCourier(m webSocketMessage) {
	var cm CourierMove

	// получить данные из m.Data и десериализовать их в структуру CourierMove

	err := json.Unmarshal(m.Data.([]byte), &cm)
	if err != nil {
		log.Fatal("websocket move data parsing error:" + err.Error())
		fmt.Println("courier_controller - m.Data parsing error:" + err.Error())

		return
	}

	// вызвать метод MoveCourier у courierService

	//fmt.Println("courier moving", cm.Direction, cm.Zoom)

	c.courierService.MoveCourier(context.Background(), cm.Direction, cm.Zoom)
}
