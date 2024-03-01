package service

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"

	"github.com/go-pg/pg/v10"
)

type UserServeRepo struct {
	rep repository.UserRepo
}

func NewUserServe(db *pg.DB) *UserServeRepo {
	return NewUserServePostgreRepo(db)
}

func NewUserServePostgreRepo(db *pg.DB) *UserServeRepo {
	//db := repository.NewPostgreDB()
	rep := repository.NewPostgreUserRepo(db)
	return &UserServeRepo{&rep}
}

func (u UserServeRepo) CreateWithArray(ctx context.Context, users []*models.User) error {
	return u.rep.CreateWithArray(ctx, users)
}

func (u UserServeRepo) GetByUsername(ctx context.Context, username string) ([]models.User, error) {
	return u.rep.GetByUsername(ctx, username)
}
func (u UserServeRepo) UpdateByUsername(ctx context.Context, username string, updateData *models.User) error {
	return u.rep.UpdateByUsername(ctx, username, updateData)
}
func (u UserServeRepo) DeleteByUsername(ctx context.Context, username string) error {
	return u.rep.DeleteByUsername(ctx, username)
}
func (u UserServeRepo) Login(ctx context.Context, username, token string) error {
	return u.rep.Login(ctx, username, token)
}
func (u UserServeRepo) Logout(ctx context.Context, token string) error {

	return u.rep.Logout(ctx, token)
}
func (u UserServeRepo) CreateUser(ctx context.Context, user *models.User) error {
	return u.rep.CreateUser(ctx, user)
}
