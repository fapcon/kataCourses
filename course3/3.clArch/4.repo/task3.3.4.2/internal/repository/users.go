package repository

import (
	"context"
	"errors"
	"petstore/internal/models"
	"reflect"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-pg/pg/v10"
)

type PostgreUserRepo struct {
	db     *pg.DB
	sqlBlb sq.StatementBuilderType
	sync.Mutex
}

func NewPostgreUserRepo(db *pg.DB) PostgreUserRepo {

	return PostgreUserRepo{
		db:     db,
		sqlBlb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (rep *PostgreUserRepo) CreateWithArray(ctx context.Context, users []*models.User) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	strInf := GetStructInfo(&models.User{})

	fields := strInf.Fields

	query := rep.sqlBlb.Insert("users").Columns(fields...)

	for _, user := range users {
		values := make([]interface{}, 0)
		strInf = GetStructInfo(user)
		for i, valueptr := range strInf.Pointers {
			if strInf.Fields[i] == "id" {
				continue
			}
			values = append(values, reflect.ValueOf(valueptr).Elem().Interface())
		}
		query = query.Values(values)
	}

	sql, _, err := query.ToSql()
	if err != nil {
		err = errors.New("query to sql error: " + err.Error())
		return err
	}

	_, err = rep.db.ExecContext(ctx, sql)
	if err != nil {
		err = errors.New("query execution error: " + err.Error())

	}

	return err
}

func (rep *PostgreUserRepo) GetByUsername(ctx context.Context, username string) ([]models.User, error) {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	var usersFound []models.User

	err := rep.db.Model(usersFound).Context(ctx).Where("username = ?", username).Select()
	if err != nil {
		err = errors.New("users getting error: " + err.Error())
	}

	return usersFound, err

}

func (rep *PostgreUserRepo) UpdateByUsername(ctx context.Context, username string, updateData *models.User) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	userMap := make(map[string]interface{})
	for i := 0; i < reflect.TypeOf(*updateData).NumField(); i++ {
		field := reflect.TypeOf(*updateData).Field(i)
		userMap[field.Name] = reflect.ValueOf(updateData).Elem().Field(i).Interface()
	}

	query, params, err := rep.sqlBlb.Update("users").SetMap(userMap).Where(sq.Eq{"username": username}).ToSql()
	if err != nil {
		return err
	}

	_, err = rep.db.ExecContext(ctx, query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (rep *PostgreUserRepo) DeleteByUsername(ctx context.Context, username string) error {

	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()
	query, params, err := rep.sqlBlb.Delete("users").Where(sq.Eq{"username": username}).ToSql()
	if err != nil {
		err = errors.New("user deletion request error: " + err.Error())
		return err
	}

	_, err = rep.db.ExecContext(ctx, query, params...)
	if err != nil {
		err = errors.New("user deletion error: " + err.Error())
		return err
	}
	return nil

}

func (rep *PostgreUserRepo) Login(ctx context.Context, username, token string) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()
	query, params, err := rep.sqlBlb.Update("users").Set("token", token).Where(sq.Eq{"username": username}).ToSql()
	if err != nil {
		err = errors.New("user login request error: " + err.Error())
		return err
	}

	_, err = rep.db.ExecContext(ctx, query, params...)
	if err != nil {
		err = errors.New("user login error: " + err.Error())
		return err
	}
	return nil

}

func (rep *PostgreUserRepo) Logout(ctx context.Context, token string) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()
	query, params, err := rep.sqlBlb.Update("users").Set("token", "").Where(sq.Eq{"token": token}).ToSql()
	if err != nil {
		err = errors.New("user logout request error: " + err.Error())
		return err
	}

	_, err = rep.db.ExecContext(ctx, query, params...)
	if err != nil {
		err = errors.New("user logout error: " + err.Error())
		return err
	}
	return nil

}

func (rep *PostgreUserRepo) CreateUser(ctx context.Context, user *models.User) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	strInf := GetStructInfo(&models.User{})

	fields := strInf.Fields

	queryRough := rep.sqlBlb.Insert("users").Columns(fields...)

	values := make([]interface{}, 0)
	strInf = GetStructInfo(user)
	for i, valueptr := range strInf.Pointers {
		if strInf.Fields[i] == "id" {
			continue
		}
		values = append(values, reflect.ValueOf(valueptr).Elem().Interface())
	}
	queryRough = queryRough.Values(values)

	query, _, err := queryRough.ToSql()
	if err != nil {
		err = errors.New("query to sql error: " + err.Error())
		return err
	}

	_, err = rep.db.ExecContext(ctx, query)
	if err != nil {
		err = errors.New("query execution error: " + err.Error())

	}

	return err
}
