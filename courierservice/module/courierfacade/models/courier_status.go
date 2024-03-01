package models

import (
	cm "courierservice/module/courier/models"
	om "courierservice/module/order/models"
)

type CourierStatus struct {
	Courier cm.Courier `json:"courier"`
	Orders  []om.Order `json:"orders"`
}
