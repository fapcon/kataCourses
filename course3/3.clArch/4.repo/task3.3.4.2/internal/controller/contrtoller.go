package controller

import "github.com/go-pg/pg/v10"

type Controller struct {
	PetCtrl PetCtrlIntf
	UsCtrl  UserCtrlIntf
	StCtrl  StoreCtrlIntf
}

func NewController(db *pg.DB) *Controller {
	return &Controller{
		PetCtrl: NewPetCntrl(db),
		UsCtrl:  NewUserCntrl(db),
		StCtrl:  NewStoreCntrl(db),
	}

}
