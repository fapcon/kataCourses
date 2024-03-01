package controller

import (
	"postgrerepository/internal/repository" //"usr/local/go/src/postgrerepository/internal/repository" //
	"postgrerepository/internal/service"    //"usr/local/go/src/postgrerepository/internal/service"    //
)

type Controller struct {
	User *service.UserController
}

func NewController(rep repository.UserRepo) *Controller {
	return &Controller{service.NewUserController(rep)}
}
