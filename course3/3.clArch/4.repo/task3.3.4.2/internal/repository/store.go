package repository

import (
	"context"
	"errors"
	"petstore/internal/models"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-pg/pg/v10"
)

type PostgreStoreRepo struct {
	db     *pg.DB
	sqlBlb sq.StatementBuilderType
	sync.Mutex
}

func NewPostgreStoreRepo(db *pg.DB) PostgreStoreRepo {

	return PostgreStoreRepo{
		db:     db,
		sqlBlb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (rep *PostgreStoreRepo) Create(ctx context.Context, order *models.Order) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	_, err := rep.db.Model(order).Context(ctx).Insert(order)

	return err
}

func (rep *PostgreStoreRepo) GetByID(ctx context.Context, id int64) (models.Order, error) {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	order := models.Order{}
	err := rep.db.Model(order).Where("id = ?", id).Select()
	if err != nil {
		err = errors.New("order getting error: " + err.Error())
	}

	return order, err
}

func (rep *PostgreStoreRepo) DeleteByID(ctx context.Context, id int64) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()
	sql, args, err := rep.sqlBlb.Delete("orders").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		err = errors.New("order deletion request error: " + err.Error())
		return err
	}

	_, err = rep.db.ExecContext(ctx, sql, args...)
	if err != nil {
		err = errors.New("order deletion error: " + err.Error())
		return err
	}
	return nil
}
