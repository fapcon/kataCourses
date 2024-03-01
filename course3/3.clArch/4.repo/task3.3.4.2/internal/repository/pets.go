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

type PostgrePetRepo struct {
	db     *pg.DB
	sqlBlb sq.StatementBuilderType
	sync.Mutex
}

func NewPostgrePetRepo(db *pg.DB) PostgrePetRepo {

	return PostgrePetRepo{
		db:     db,
		sqlBlb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (rep *PostgrePetRepo) Create(ctx context.Context, pet *models.Pet) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	_, err := rep.db.Model(pet).Context(ctx).Insert(pet)

	return err
}

func (rep *PostgrePetRepo) Update(ctx context.Context, pet *models.Pet) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	petMap := make(map[string]interface{})
	for i := 0; i < reflect.TypeOf(*pet).NumField(); i++ {
		field := reflect.TypeOf(*pet).Field(i)
		petMap[field.Name] = reflect.ValueOf(pet).Elem().Field(i).Interface()
	}

	sql, values, err := rep.sqlBlb.Update(pet.TableName()).SetMap(petMap).Where(sq.Eq{"id": pet.ID}).ToSql()
	if err != nil {
		return err
	}

	_, err = rep.db.ExecContext(ctx, sql, values)
	if err != nil {
		return err
	}

	return nil
}

func (rep *PostgrePetRepo) GetByStatus(ctx context.Context, status string) ([]models.Pet, error) {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	pets := []models.Pet{}
	err := rep.db.Model(pets).Where("status = ?", status).Select()
	if err != nil {
		err = errors.New("pet getting error: " + err.Error())
	}

	return pets, err
}

func (rep *PostgrePetRepo) GetByID(ctx context.Context, id int64) (models.Pet, error) {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	pet := models.Pet{}
	err := rep.db.Model(pet).Where("id = ?", id).Select()
	if err != nil {
		err = errors.New("pet getting error: " + err.Error())
	}

	return pet, err
}

func (rep *PostgrePetRepo) UpdateByID(ctx context.Context, id int64, name, status string) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	sql, values, err := rep.sqlBlb.Update("pets").Set("name", name).Set("status", status).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		err = errors.New("pet updating request error: " + err.Error())
		return err
	}

	_, err = rep.db.ExecContext(ctx, sql, values...)
	if err != nil {
		err = errors.New("pet updating error: " + err.Error())
		return err
	}

	return nil
}

func (rep *PostgrePetRepo) DeleteByID(ctx context.Context, id int64) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()
	sql, args, err := rep.sqlBlb.Delete("pets").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		err = errors.New("pet deletion request error: " + err.Error())
		return err
	}

	_, err = rep.db.ExecContext(ctx, sql, args...)
	if err != nil {
		err = errors.New("pet deletion error: " + err.Error())
		return err
	}
	return nil
}
