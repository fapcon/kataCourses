package service

import (
	"context"
	cservice "courierservice/module/courier/service"
	cfm "courierservice/module/courierfacade/models"
	oservice "courierservice/module/order/service"
	"fmt"
	"log"
)

const (
	CourierVisibilityRadius = 2800 // 2.8km
)

type CourierFacer interface {
	MoveCourier(ctx context.Context, direction, zoom int) // отвечает за движение курьера по карте direction - направление движения, zoom - уровень зума
	GetStatus(ctx context.Context) cfm.CourierStatus      // отвечает за получение статуса курьера и заказов вокруг него
}

// CourierFacade фасад для курьера и заказов вокруг него (для фронта)
type CourierFacade struct {
	courierService cservice.Courierer
	orderService   oservice.Orderer
}

func (cf CourierFacade) MoveCourier(ctx context.Context, direction, zoom int) {
	//fmt.Println("facade - move courier entered")

	courier, err := cf.courierService.GetCourier(ctx)

	if err != nil {
		log.Fatal("courier getting error:" + err.Error())
		return
	}

	err = cf.courierService.MoveCourier(*courier, direction, zoom)
	if err != nil {
		log.Fatal("courier moving error:" + err.Error())

	}
	//fmt.Println("facade - move courier finished")
}

func (cf CourierFacade) GetStatus(ctx context.Context) cfm.CourierStatus {
	//fmt.Println("facade - get status entered")
	courier, err := cf.courierService.GetCourier(ctx)

	if err != nil {
		log.Fatal("courier getting error:" + err.Error())
		return cfm.CourierStatus{}
	}

	orders, err := cf.orderService.GetByRadius(ctx, courier.Location.Lng, courier.Location.Lat, CourierVisibilityRadius, "m")
	if err != nil {
		log.Fatal("orders getting error:" + err.Error())
		fmt.Println("orders getting error:" + err.Error())
		return cfm.CourierStatus{}
	}
	//fmt.Println("facade - get status, found orders nearby:", len(orders))
	return cfm.CourierStatus{
		Courier: *courier,
		Orders:  orders}
}

func NewCourierFacade(courierService cservice.Courierer, orderService oservice.Orderer) CourierFacer {
	return &CourierFacade{courierService: courierService, orderService: orderService}
}
