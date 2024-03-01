package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"courierservice/module/order/models"

	"github.com/go-redis/redis/v8"
)

type OrderStorager interface {
	Save(ctx context.Context, order models.Order, maxAge time.Duration) error                       // сохранить заказ с временем жизни
	GetByID(ctx context.Context, orderID int) (*models.Order, error)                                // получить заказ по id
	GenerateUniqueID(ctx context.Context) (int64, error)                                            // сгенерировать уникальный id
	GetByRadius(ctx context.Context, lng, lat, radius float64, unit string) ([]models.Order, error) // получить заказы в радиусе от точки
	GetCount(ctx context.Context) (int, error)                                                      // получить количество заказов
	RemoveOldOrders(ctx context.Context, maxAge time.Duration) error                                // удалить старые заказы по истечению времени maxAge
}

type OrderStorage struct {
	storage *redis.Client
}

func NewOrderStorage(storage *redis.Client) OrderStorager {
	return &OrderStorage{storage: storage}
}

func (o *OrderStorage) Save(ctx context.Context, order models.Order, maxAge time.Duration) error {
	// save with geogrpc redis

	return o.saveOrderWithGeo(ctx, order, maxAge)
}

func (o *OrderStorage) RemoveOldOrders(ctx context.Context, maxAge time.Duration) error {
	// получить ID всех старых ордеров, которые нужно удалить
	// используя метод ZRangeByScore
	// старые ордеры это те, которые были созданы две минуты назад
	// и более
	/**
	&redis.ZRangeBy{
		Max: использовать вычисление времени с помощью maxAge,
		Min: "0",
	}
	*/

	t := time.Now().Add(-maxAge).Unix()
	ts := strconv.Itoa(int(t))
	orders, err := o.storage.ZRangeByScore(ctx, "orders", &redis.ZRangeBy{
		Max: ts,
		Min: "0",
	}).Result()

	if err != nil {
		return err
	}

	// Проверить количество старых ордеров

	if len(orders) == 0 {
		return err
	}
	// удалить старые ордеры из redis используя метод ZRemRangeByScore где ключ "orders" min "-inf" max "(время создания старого ордера)"
	// удалять ордера по ключу не нужно, они будут удалены автоматически по истечению времени жизни
	o.storage.ZRemRangeByScore(ctx, "orders", "-inf", ts)
	o.storage.ZRemRangeByScore(ctx, "orders_locations", "-inf", ts)
	return err
}

func (o *OrderStorage) GetByID(ctx context.Context, orderID int) (*models.Order, error) {

	order := &models.Order{}
	// получаем ордер из redis по ключу order:ID

	// проверяем что ордер не найден исключение redis.Nil, в этом случае возвращаем nil, nil
	orderJSON := o.storage.Get(ctx, "orders:"+strconv.Itoa(orderID))

	if err := orderJSON.Err(); err == redis.Nil {
		return nil, err
	}
	// десериализуем ордер из json
	err := json.Unmarshal([]byte(orderJSON.String()), order)
	if err != nil {
		log.Fatal("order parsing error:" + err.Error())
		return order, err
	}

	return order, nil
}

func (o *OrderStorage) saveOrderWithGeo(ctx context.Context, order models.Order, maxAge time.Duration) error {

	// сериализуем ордер в json
	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Fatal("order parsing error:" + err.Error())
		fmt.Println("order parsing error:", err)
		return err
	}

	// сохраняем ордер в json redis по ключу order:ID с временем жизни maxAge
	orderKey := "orders:" + strconv.Itoa(int(order.ID))
	err = o.storage.Set(ctx, orderKey, orderJSON, maxAge).Err()
	if err != nil {
		log.Fatal("order saving error:" + err.Error())
		fmt.Println("order saving error:", err)
		return err
	}

	// добавляем ордер в гео индекс используя метод GeoAdd где Name - это ключ ордера, а Longitude и Latitude - координаты
	err = o.storage.GeoAdd(ctx, "orders_locations", &redis.GeoLocation{
		Name:      orderKey,
		Longitude: order.Lng,
		Latitude:  order.Lat,
	}).Err()

	if err != nil {
		log.Fatal("location saving error:" + err.Error())
		fmt.Println("location saving error:", err)
		return err
	}
	//fmt.Println("order location added by GeoAdd for:", orderKey, order.Lng, order.Lat)
	// zset сохраняем ордер для получения количества заказов со сложностью O(1)
	//res, err := o.getOrdersByRadius(ctx, order.Lng, order.Lat, 2800, "m")

	//fmt.Println("getting and check orders by radius for it:", res, err)

	err = o.storage.ZAdd(ctx, "orders", &redis.Z{
		Score:  float64(order.CreatedAt.Unix()),
		Member: orderKey,
	}).Err()
	if err != nil {
		log.Fatal("order score saving error:" + err.Error())
		fmt.Println("order score saving error:", err)
		return err
	}
	//fmt.Println("order score added by ZAdd for:", orderKey)

	// Score - время создания ордера

	return err
}

func (o *OrderStorage) GetCount(ctx context.Context) (int, error) {
	// получить количество ордеров в упорядоченном множестве используя метод ZCard
	count, err := o.storage.ZCard(ctx, "orders").Result()
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (o *OrderStorage) GetByRadius(ctx context.Context, lng, lat, radius float64, unit string) ([]models.Order, error) {

	// используем метод getOrdersByRadius для получения ID заказов в радиусе
	ordersLocation, err := o.getOrdersByRadius(ctx, lng, lat, radius, unit)
	// обратите внимание, что в случае отсутствия заказов в радиусе
	// метод getOrdersByRadius должен вернуть nil, nil (при ошибке redis.Nil)
	//fmt.Println("order storage found orders nearby:", len(ordersLocation))
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	orders := make([]models.Order, len(ordersLocation))
	// проходим по списку ID заказов и получаем данные о заказе
	// получаем данные о заказе по ID из redis по ключу order:ID
	for i, orderLocation := range ordersLocation {
		orderJSON := o.storage.Get(ctx, orderLocation.Name)
		err = orderJSON.Err()
		if err != nil {
			log.Fatal("order location request error:" + err.Error())
			return orders, err
		}

		order := &models.Order{}
		err = json.Unmarshal([]byte(orderJSON.Val()), order)
		if err != nil {
			fmt.Println("orderJSON.Val:", orderJSON.Val())
			log.Fatal("order location parsing error:" + err.Error())
			return orders, err
		}

		orders[i] = *order

	}

	return orders, nil
}

func (o *OrderStorage) getOrdersByRadius(ctx context.Context, lng, lat, radius float64, unit string) ([]redis.GeoLocation, error) {
	// в данном методе мы получаем список ордеров в радиусе от точки
	// возвращаем список ордеров с координатами и расстоянием до точки

	query := &redis.GeoRadiusQuery{
		Radius:    radius,
		Unit:      unit,
		WithCoord: true,
		WithDist:  true,
		//WithGeoHash: true,
	}

	locations, err := o.storage.GeoRadius(ctx, "orders_locations", lng, lat, query).Result()
	if err != nil {
		if err == redis.Nil {
			//fmt.Println("no locations returned from order_storage")
			err = nil
		} else {
			log.Fatal("order location request error:" + err.Error())
			fmt.Println("order location request error:" + err.Error())
		}

	}

	//fmt.Println("order locations returned:", len(locations))
	return locations, err

}

func (o *OrderStorage) GenerateUniqueID(ctx context.Context) (int64, error) {
	var err error
	var id int64

	// генерируем уникальный ID для ордера
	// используем для этого redis incr по ключу order:id

	id, err = o.storage.Incr(ctx, "orders:id").Result()

	return id, err
}
