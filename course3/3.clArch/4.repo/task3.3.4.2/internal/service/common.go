package service

import "github.com/go-pg/pg/v10"

type Serve struct {
	PetServe
	UserServe
	StoreServe
}

func NewServe(db *pg.DB) *Serve {
	return &Serve{
		NewPetServe(db),
		NewUserServe(db),
		NewStoreServe(db),
	}
}
