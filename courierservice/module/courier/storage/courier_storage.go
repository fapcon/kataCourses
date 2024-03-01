package storage

import (
	"context"
	"encoding/json"
	"log"

	"courierservice/module/courier/models"

	"github.com/go-redis/redis/v8"
)

type CourierStorager interface {
	Save(ctx context.Context, courier models.Courier) error // сохранить курьера по ключу courier
	GetOne(ctx context.Context) (*models.Courier, error)    // получить курьера по ключу courier
}

type CourierStorage struct {
	storage *redis.Client
}

func NewCourierStorage(storage *redis.Client) CourierStorager {
	return &CourierStorage{storage: storage}
}

func (cs CourierStorage) Save(ctx context.Context, courier models.Courier) error {
	courierJSON, err := json.Marshal(courier)
	if err != nil {
		log.Fatal("json parsing error:" + err.Error())
		return err
	}
	err = cs.storage.Set(ctx, "courier", courierJSON, 0).Err()
	if err != nil {
		log.Fatal("storage set-request error:" + err.Error())
		return err
	}

	return err

}

func (cs CourierStorage) GetOne(ctx context.Context) (*models.Courier, error) {
	courier := &models.Courier{}
	courierJSON, err := cs.storage.Get(ctx, "courier").Bytes()
	if err != nil {
		log.Fatal("storage get-request error:" + err.Error())
		return courier, err
	}

	err = json.Unmarshal(courierJSON, courier)
	if err != nil {
		log.Fatal("storage respond parsing error:" + err.Error())

	}
	return courier, err

}
