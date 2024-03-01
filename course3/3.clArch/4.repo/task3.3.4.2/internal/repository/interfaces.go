package repository

import (
	"context"
	"petstore/internal/models"
)

type Repo struct {
	PetRepo
	UserRepo
	StoreRepo
}

func NewRepo() *Repo {
	return &Repo{}
}

type PetRepo interface {
	Create(ctx context.Context, pet *models.Pet) error
	Update(ctx context.Context, pet *models.Pet) error
	GetByStatus(ctx context.Context, status string) ([]models.Pet, error)
	GetByID(ctx context.Context, id int64) (models.Pet, error)
	UpdateByID(ctx context.Context, id int64, name, status string) error
	DeleteByID(ctx context.Context, id int64) error
}

type StoreRepo interface {
	Create(ctx context.Context, order *models.Order) error
	GetByID(ctx context.Context, id int64) (models.Order, error)
	DeleteByID(ctx context.Context, id int64) error
}

type UserRepo interface {
	CreateWithArray(ctx context.Context, users []*models.User) error
	GetByUsername(ctx context.Context, username string) ([]models.User, error)
	UpdateByUsername(ctx context.Context, username string, updateData *models.User) error
	DeleteByUsername(ctx context.Context, username string) error
	Login(ctx context.Context, username, token string) error
	Logout(ctx context.Context, token string) error
	CreateUser(ctx context.Context, users *models.User) error
}
