package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"postgrerepository/internal/models" //"usr/local/go/src/postgrerepository/internal/models" //
	"strconv"
	"strings"
)

type UserController struct {
	userRepo UserRepo
}

type UserRepo interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, user *models.User, conditions *models.Conditions) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, c *models.Conditions) (*[]models.User, error)
}

func NewUserController(usrep UserRepo) *UserController {
	return &UserController{userRepo: usrep}
}

func (uscont *UserController) Create(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&models.User{})
	if err != nil {
		http.Error(w, "json decoding error:"+err.Error(), http.StatusBadRequest)
	}

	fmt.Fprint(w, "user created", http.StatusOK)
}

func (uscont *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	pathSpl := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathSpl[len(pathSpl)-1])

	if err != nil {
		http.Error(w, "ID error:"+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := uscont.userRepo.GetByID(context.Background(), id)

	if err != nil {
		http.Error(w, "user search error:"+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, user, http.StatusOK)
}

func (uscont *UserController) Update(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, "user request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}
	conds := &models.Conditions{}
	err = json.NewDecoder(r.Body).Decode(conds)
	if err != nil {
		http.Error(w, "request condition parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = uscont.userRepo.Update(context.Background(), user, conds)

	if err != nil {
		http.Error(w, "user updating error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "user updated")
}

func (uscont *UserController) Delete(w http.ResponseWriter, r *http.Request) {

	pathSpl := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathSpl[len(pathSpl)-1])

	if err != nil {
		http.Error(w, "ID error:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = uscont.userRepo.Delete(context.Background(), id)

	if err != nil {
		http.Error(w, "user deletion error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "user deleted")

}

func (uscont *UserController) List(w http.ResponseWriter, r *http.Request) {
	cond := &models.Conditions{}
	err := json.NewDecoder(r.Body).Decode(cond)
	if err != nil {
		http.Error(w, "conditions parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	users, err := uscont.userRepo.List(context.Background(), cond)
	if err != nil {
		http.Error(w, "users getting error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, users)
}
