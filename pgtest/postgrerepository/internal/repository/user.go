package repository

import (
	"context"
	"errors"
	"postgrerepository/internal/models" //"usr/local/go/src/postgrerepository/internal/models" //
	"reflect"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-pg/pg/v10"
)

type UserRepo interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, user *models.User, conditions *models.Conditions) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, c *models.Conditions) (*[]models.User, error)
}

type PostgreUserRepo struct {
	db     *pg.DB
	sqlBld sq.StatementBuilderType
}

func NewPostgreUserRepo(db *pg.DB) PostgreUserRepo {
	return PostgreUserRepo{
		db:     db,
		sqlBld: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (rep PostgreUserRepo) Create(ctx context.Context, user *models.User) error {
	_, err := rep.db.Model(user).Context(ctx).Insert(user)
	return err

}

func (rep PostgreUserRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}
	err := rep.db.Model(user).Where("id = ?", id).Select()
	if user.Deleted {
		user = &models.User{}
		err = errors.New("user deleted")
	}

	return user, err

}

func (rep PostgreUserRepo) Update(ctx context.Context, user *models.User, cond *models.Conditions) error {
	userMap := make(map[string]interface{})
	for i := 0; i < reflect.TypeOf(*user).NumField(); i++ {
		field := reflect.TypeOf(*user).Field(i)
		userMap[field.Name] = reflect.ValueOf(user).Elem().Field(i).Interface()
	}

	condMap := make(map[string]interface{})

	for key, vl := range cond.Equal {
		condMap[key] = vl
	}

	query, vls, err := rep.sqlBld.Update("users").SetMap(userMap).Where(condMap).ToSql()
	if err != nil {
		return errors.New("query building error:" + err.Error())
	}

	_, err = rep.db.ExecContext(ctx, query, vls)

	if err != nil {
		err = errors.New("query execution error:" + err.Error())
	}

	return err

}

func (rep PostgreUserRepo) Delete(ctx context.Context, id int) error {
	user := &models.User{}

	err := rep.db.Model(user).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return errors.New("user deletion error:" + err.Error())
	}

	user.Deleted = true
	_, err = rep.db.Model(user).Where("id = ?", id).Update()

	return err

}

func (rep PostgreUserRepo) List(ctx context.Context, cond *models.Conditions) (*[]models.User, error) {
	users := make([]models.User, 0)

	query := rep.db.Model(&users).Context(ctx)

	for key, vl := range cond.Equal {
		query = query.Where(key+" = ?", vl)
	}

	for key, vl := range cond.NotEqual {
		query = query.Where(key+" != ?", vl)
	}

	for _, ord := range cond.Order {
		if ord.Asc {
			query = query.Order(ord.Field + " ASC")
		} else {
			query = query.Order(ord.Field + " DESC")
		}
	}

	if cond.LimitOffset != nil {
		query = query.Limit(int(cond.LimitOffset.Limit)).Offset(int(cond.LimitOffset.Offset))
	}
	err := query.Select()
	if err != nil {
		err = errors.New("user list query error:" + err.Error())
	}

	if len(users) == 0 {
		err = errors.New("no users match to request")
	}

	return &users, err

}
