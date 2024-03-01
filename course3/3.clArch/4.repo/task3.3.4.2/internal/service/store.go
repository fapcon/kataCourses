package service

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"

	"github.com/go-pg/pg/v10"
)

type StoreServeRepo struct {
	rep repository.StoreRepo
}

func NewStoreServe(db *pg.DB) *StoreServeRepo {
	return NewStoreServePostgreRepo(db)
}

func NewStoreServePostgreRepo(db *pg.DB) *StoreServeRepo {
	//db := repository.NewPostgreDB()
	rep := repository.NewPostgreStoreRepo(db)
	return &StoreServeRepo{&rep}
}

func (s StoreServeRepo) Create(ctx context.Context, order *models.Order) error {
	return s.rep.Create(ctx, order)
}

func (s StoreServeRepo) GetByID(ctx context.Context, id int64) (models.Order, error) {
	return s.rep.GetByID(ctx, id)
}

func (s StoreServeRepo) DeleteByID(ctx context.Context, id int64) error {
	return s.rep.DeleteByID(ctx, id)
}
