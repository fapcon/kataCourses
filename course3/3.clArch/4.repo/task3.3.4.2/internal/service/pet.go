package service

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"

	"github.com/go-pg/pg/v10"
)

type PetServeRepo struct {
	rep repository.PetRepo
}

func NewPetServe(db *pg.DB) *PetServeRepo {
	return NewPetServePostgreRepo(db)
}

func NewPetServePostgreRepo(db *pg.DB) *PetServeRepo {
	//db := repository.NewPostgreDB()
	rep := repository.NewPostgrePetRepo(db)
	return &PetServeRepo{&rep}
}

func (p PetServeRepo) Create(ctx context.Context, pet *models.Pet) error {
	return p.rep.Create(ctx, pet)
}
func (p PetServeRepo) Update(ctx context.Context, pet *models.Pet) error {
	return p.rep.Update(ctx, pet)
}
func (p PetServeRepo) GetByStatus(ctx context.Context, status string) ([]models.Pet, error) {
	return p.rep.GetByStatus(ctx, status)
}
func (p PetServeRepo) GetByID(ctx context.Context, id int64) (models.Pet, error) {
	return p.rep.GetByID(ctx, id)
}
func (p PetServeRepo) UpdateByID(ctx context.Context, id int64, name, status string) error {
	return p.rep.UpdateByID(ctx, id, name, status)
}
func (p PetServeRepo) DeleteByID(ctx context.Context, id int64) error {
	return p.rep.DeleteByID(ctx, id)
}
