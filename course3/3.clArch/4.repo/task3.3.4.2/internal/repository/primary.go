package repository

import (
	"github.com/go-pg/pg/v10"
)

type PostgreRepo struct {
	pet   PostgrePetRepo
	user  PostgreUserRepo
	store PostgreStoreRepo
}

func NewPostgreRepo(db *pg.DB) PostgreRepo {

	return PostgreRepo{
		pet:   NewPostgrePetRepo(db),
		user:  NewPostgreUserRepo(db),
		store: NewPostgreStoreRepo(db),
	}
}
