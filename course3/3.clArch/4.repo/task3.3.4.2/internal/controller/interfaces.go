package controller

import "net/http"

type PetCtrlIntf interface {
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	GetByStatus(http.ResponseWriter, *http.Request)
	GetById(http.ResponseWriter, *http.Request)
	UpdateById(http.ResponseWriter, *http.Request)
	DeleteById(http.ResponseWriter, *http.Request)
}

type UserCtrlIntf interface {
	CreateWithArray(http.ResponseWriter, *http.Request)
	GetByUsername(http.ResponseWriter, *http.Request)
	UpdateByUsername(http.ResponseWriter, *http.Request)
	DeleteByUsername(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	Logout(http.ResponseWriter, *http.Request)
	CreateUser(http.ResponseWriter, *http.Request)
}

type StoreCtrlIntf interface {
	Create(http.ResponseWriter, *http.Request)
	GetByID(http.ResponseWriter, *http.Request)
	DeleteByID(http.ResponseWriter, *http.Request)
}
