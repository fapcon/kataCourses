package service

import (
	"awesomeProject/internal/models/user"
	"awesomeProject/internal/repository"
	"context"
	"errors"
	"time"
)

type UserServiceI interface {
	CreateUser(ctx context.Context, u *user.User) error
	GetUserByID(ctx context.Context, id string) (*user.User, error)
	UpdateUser(ctx context.Context, u *user.User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, limit, offset int) ([]user.User, error)
}

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, u *user.User) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return s.repo.Create(ctx, *u)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	r, err := s.repo.GetByID(ctx, id)
	return &r, err
}

func (s *UserService) UpdateUser(ctx context.Context, u *user.User) error {
	existingUser, err := s.repo.GetByID(ctx, u.ID)
	if err != nil {
		return err
	}
	if existingUser.ID == "" {
		return errors.New("user not found")
	}

	u.UpdatedAt = time.Now()

	return s.repo.Update(ctx, *u)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context, limit, offset int) ([]user.User, error) {
	return s.repo.List(ctx, limit, offset)
}
