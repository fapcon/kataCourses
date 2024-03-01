package service

import (
	"context"
	"courierservice/geo"
	"courierservice/module/courier/models"
	"courierservice/module/courier/storage"
	"log"
	"math"
)

// Направления движения курьера
const (
	DirectionUp    = 0
	DirectionDown  = 1
	DirectionLeft  = 2
	DirectionRight = 3
)

const (
	DefaultCourierLat = 59.9311
	DefaultCourierLng = 30.3609
)

type Courierer interface {
	GetCourier(ctx context.Context) (*models.Courier, error)
	MoveCourier(courier models.Courier, direction, zoom int) error
}

type CourierService struct {
	courierStorage storage.CourierStorager
	allowedZone    geo.PolygonChecker
	disabledZones  []geo.PolygonChecker
}

func NewCourierService(courierStorage storage.CourierStorager, allowedZone geo.PolygonChecker, disabledZones []geo.PolygonChecker) Courierer {
	courier := models.Courier{}
	//fmt.Println("courier service - getting random location...")
	courier.Location = models.Point(geo.GetRandomAllowedLocation(allowedZone, disabledZones))
	//fmt.Println("courier service - location got")
	//time.Sleep(1 * time.Second)
	courierStorage.Save(context.Background(), courier)
	//fmt.Println("courier service - new courier saved")

	return &CourierService{courierStorage: courierStorage, allowedZone: allowedZone, disabledZones: disabledZones}
}

func (c *CourierService) GetCourier(ctx context.Context) (*models.Courier, error) {
	// получаем курьера из хранилища используя метод GetOne из storage/courier.go
	courier, err := c.courierStorage.GetOne(ctx)
	if err != nil {
		log.Fatal("courier getting error:" + err.Error())
		return courier, err
	}

	// проверяем, что курьер находится в разрешенной зоне
	// если нет, то перемещаем его в случайную точку в разрешенной зоне

	if !geo.CheckPointIsAllowed(geo.Point(courier.Location), c.allowedZone, c.disabledZones) {
		newLocation := geo.GetRandomAllowedLocation(c.allowedZone, c.disabledZones)
		courier.Location = models.Point(newLocation)
		// сохраняем новые координаты курьера
		err = c.courierStorage.Save(ctx, *courier)
		if err != nil {
			log.Fatal("courier saving error:" + err.Error())
			return courier, err
		}

	}

	return courier, err
}

// MoveCourier : direction - направление движения курьера, zoom - зум карты
func (c *CourierService) MoveCourier(courier models.Courier, direction, zoom int) error {
	// точность перемещения зависит от зума карты использовать формулу 0.001 / 2^(zoom - 14)
	// 14 - это максимальный зум карты

	movement := 0.001 / math.Pow(2, float64(zoom-14))

	switch direction {
	case 0:
		courier.Location.Lat += movement
	case 1:
		courier.Location.Lat -= movement
	case 2:
		courier.Location.Lng -= movement
	case 3:
		courier.Location.Lng += movement
	}

	// далее нужно проверить, что курьер не вышел за границы зоны
	// если вышел, то нужно переместить его в случайную точку внутри зоны

	if !geo.CheckPointIsAllowed(geo.Point(courier.Location), c.allowedZone, c.disabledZones) {
		courier.Location = models.Point(geo.GetRandomAllowedLocation(c.allowedZone, c.disabledZones))
	}

	// далее сохранить изменения в хранилище

	err := c.courierStorage.Save(context.Background(), courier)

	if err != nil {
		log.Fatal("courier saving error" + err.Error())
	}

	return err
}
