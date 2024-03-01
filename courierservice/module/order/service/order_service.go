package service

import (
	"context"
	"courierservice/geo"
	"courierservice/module/order/models"
	"courierservice/module/order/storage"
	"log"
	"math/rand"
	"time"
)

const (
	minDeliveryPrice = 100.00
	maxDeliveryPrice = 500.00

	maxOrderPrice = 3000.00
	minOrderPrice = 1000.00

	orderMaxAge = 2 * time.Minute
)

type Orderer interface {
	GetByRadius(ctx context.Context, lng, lat, radius float64, unit string) ([]models.Order, error) // возвращает заказы через метод storage.GetByRadius
	Save(ctx context.Context, order models.Order) error                                             // сохраняет заказ через метод storage.Save с заданным временем жизни OrderMaxAge
	GetCount(ctx context.Context) (int, error)                                                      // возвращает количество заказов через метод storage.GetCount
	RemoveOldOrders(ctx context.Context) error                                                      // удаляет старые заказы через метод storage.RemoveOldOrders с заданным временем жизни OrderMaxAge
	GenerateOrder(ctx context.Context) error                                                        // генерирует заказ в случайной точке из разрешенной зоны, с уникальным id, ценой и ценой доставки
}

// OrderService реализация интерфейса Orderer
// в нем должны быть методы GetByRadius, Save, GetCount, RemoveOldOrders, GenerateOrder
// данный сервис отвечает за работу с заказами
type OrderService struct {
	storage       storage.OrderStorager
	allowedZone   geo.PolygonChecker
	disabledZones []geo.PolygonChecker
}

func (os OrderService) GetByRadius(ctx context.Context, lng, lat, radius float64, unit string) ([]models.Order, error) {
	// возвращает заказы через метод storage.GetByRadius
	return os.storage.GetByRadius(ctx, lng, lat, radius, unit)
}
func (os OrderService) Save(ctx context.Context, order models.Order) error {
	// сохраняет заказ через метод storage.Save с заданным временем жизни OrderMaxAge
	return os.storage.Save(ctx, order, orderMaxAge)
}
func (os OrderService) GetCount(ctx context.Context) (int, error) {
	// возвращает количество заказов через метод storage.GetCount
	return os.storage.GetCount(ctx)
}
func (os OrderService) RemoveOldOrders(ctx context.Context) error {
	// удаляет старые заказы через метод storage.RemoveOldOrders с заданным временем жизни OrderMaxAge
	return os.storage.RemoveOldOrders(ctx, orderMaxAge)
}
func (os OrderService) GenerateOrder(ctx context.Context) error {
	id, err := os.storage.GenerateUniqueID(ctx)
	if err != nil {
		log.Fatal("order generating error:" + err.Error())
	}

	point := geo.GetRandomAllowedLocation(os.allowedZone, os.disabledZones)

	order := models.Order{
		ID:            id,
		Price:         rand.Float64()*(maxOrderPrice-minOrderPrice) + minOrderPrice,
		DeliveryPrice: rand.Float64()*(maxDeliveryPrice-minDeliveryPrice) + minDeliveryPrice,
		Lng:           point.Lng,
		Lat:           point.Lat,
		IsDelivered:   false,
		CreatedAt:     time.Now(),
	}

	//fmt.Println("order generated:", order, point)

	return os.Save(ctx, order)
}
func NewOrderService(storage storage.OrderStorager, allowedZone geo.PolygonChecker, disallowedZone []geo.PolygonChecker) Orderer {
	return &OrderService{storage: storage, allowedZone: allowedZone, disabledZones: disallowedZone}
}
