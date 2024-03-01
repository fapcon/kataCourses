package dao

import (
	"context"
	"fmt"
	"studentgit.kata.academy/fcons/go-kata/course3/1.reflection/1.reflection_using/task3.1.1.4/tabler"

	"reflect"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=./sql_adapter.go -destination=../../mock/adapter_mock.go -package=mock
type IfaceDAO interface {
	BuildSelect(tableName string, condition Condition, fields ...string) (string, []interface{}, error)
	Create(ctx context.Context, entity tabler.Tabler, opts ...interface{}) error
	List(ctx context.Context, dest interface{}, table tabler.Tabler, condition Condition, opts ...interface{}) error
	Update(ctx context.Context, entity tabler.Tabler, condition Condition, opts ...interface{}) error
}

type Condition struct {
	Equal       map[string]interface{}
	NotEqual    map[string]interface{}
	Order       []*Order
	LimitOffset *LimitOffset
	ForUpdate   bool
	Upsert      bool
}

type Order struct {
	Field string
	Asc   bool
}

type LimitOffset struct {
	Offset int64
	Limit  int64
}

type DAO struct {
	db         *sqlx.DB
	sqlBuilder sq.StatementBuilderType
}

func NewDAO(db *sqlx.DB) IfaceDAO {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &DAO{db: db, sqlBuilder: builder}
}

func (d *DAO) BuildSelect(tableName string, condition Condition, fields ...string) (string, []interface{}, error) {
	queryRaw := d.sqlBuilder.Select(fields...).From(tableName)

	for key, val := range condition.Equal {
		queryRaw = queryRaw.Where(sq.Eq{key: val})
	}
	for key, val := range condition.NotEqual {
		queryRaw = queryRaw.Where(sq.NotEq{key: val})
	}
	for _, ord := range condition.Order {
		if ord.Asc {
			queryRaw = queryRaw.OrderBy(fmt.Sprintf("%s ASC", ord.Field))
		} else {
			queryRaw = queryRaw.OrderBy(fmt.Sprintf("%s DESC", ord.Field))
		}
	}
	if condition.LimitOffset != nil {
		queryRaw = queryRaw.Limit(uint64(condition.LimitOffset.Limit)).
			Offset(uint64(condition.LimitOffset.Offset))
	}
	return queryRaw.ToSql()
}

func filterByTag(tag string, tvalue string) func(fields *[]reflect.StructField) {
	return tabler.FilterByTags(map[string]func(value string) bool{
		tag: func(value string) bool {
			return strings.Contains(value, tvalue)
		},
	})
}

func (d *DAO) Create(ctx context.Context, entity tabler.Tabler, opts ...interface{}) error {
	var options []func(*[]reflect.StructField)
	for _, opt := range opts {
		if option, ok := opt.(func(*[]reflect.StructField)); ok {
			options = append(options, option)
		}
	}

	ins := d.sqlBuilder.Insert(entity.TableName()).Columns()
	gg := tabler.GetStructInfo(entity, options...)
	val := make([]interface{}, len(gg.Fields))
	for i, pointer := range gg.Pointers {
		val[i] = reflect.ValueOf(pointer).Elem().Interface()
	}
	sql, _, err := ins.Columns(gg.Fields...).Values(val...).ToSql()
	_, err = d.db.ExecContext(ctx, sql)
	if err != nil {
		return err
	}

	return err
}

func (d *DAO) List(ctx context.Context, dest interface{}, table tabler.Tabler, condition Condition, opts ...interface{}) error {
	var options []func(*[]reflect.StructField)
	for _, opt := range opts {
		if option, ok := opt.(func(*[]reflect.StructField)); ok {
			options = append(options, option)
		}
	}
	gg := tabler.GetStructInfo(table, options...)
	query, args, err := d.BuildSelect(table.TableName(), condition, gg.Fields...)
	if err != nil {
		return err
	}
	err = d.db.SelectContext(ctx, dest, query, args...)
	return err
}

func (d *DAO) Update(ctx context.Context, entity tabler.Tabler, condition Condition, opts ...interface{}) error {
	var options []func(*[]reflect.StructField)
	for _, opt := range opts {
		if option, ok := opt.(func(*[]reflect.StructField)); ok {
			options = append(options, option)
		}
	}
	upd := d.sqlBuilder.Update(entity.TableName())
	gg := tabler.GetStructInfo(entity, options...)
	vals := make(map[string]interface{}, len(gg.Fields))
	for i, pointer := range gg.Pointers {
		vals[gg.Fields[i]] = reflect.ValueOf(pointer).Elem().Interface()
	}
	sql, _, err := upd.SetMap(vals).Where(condition.Equal).ToSql()
	if err != nil {
		return err
	}
	_, err = d.db.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	return err
}
