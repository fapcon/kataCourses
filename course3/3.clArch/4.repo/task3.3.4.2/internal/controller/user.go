package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"petstore/internal/models"
	"petstore/internal/service"
	"strings"

	"github.com/go-chi/jwtauth"
	"github.com/go-pg/pg/v10"
)

type UserCntrl struct {
	us service.UserServe
}

func NewUserCntrl(db *pg.DB) *UserCntrl {
	return &UserCntrl{us: service.NewUserServe(db)}
}

// CreateWithArray creates multiple users from the provided array.
// @Summary Create multiple users
// @Description Create multiple users from a list
// @Tags users
// @Accept json
// @Produce json
// @Param users body []models.User true "Array of users"
// @Success 200 {string} string "Users created"
// @Failure 400 {string} string "request parsing error"
// @Failure 500 {string} string "user array creation error"
// @Router /users/createWithArray [post]
func (uc UserCntrl) CreateWithArray(w http.ResponseWriter, r *http.Request) {
	users := make([]*models.User, 0)
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = uc.us.CreateWithArray(context.Background(), users)
	if err != nil {
		http.Error(w, "user array creation error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Users created")

}

// UpdateByUsername updates a user's information based on their username.
// @Summary Update a user by username
// @Description Update a user's information by their username
// @Tags users
// @Accept json
// @Produce json
// @Param username path string true "Username of the user to update"
// @Param user body models.User true "Updated user information"
// @Success 200 {string} string "Users updated"
// @Failure 400 {string} string "request parsing error"
// @Failure 417 {string} string "user update error"
// @Router /users/{username} [put]
func (uc UserCntrl) GetByUsername(w http.ResponseWriter, r *http.Request) {
	pathSpl := strings.Split(r.URL.Path, "/")
	username := pathSpl[len(pathSpl)-1]

	if username == "" {
		http.Error(w, "empty name error", http.StatusBadRequest)
		return
	}

	user, err := uc.us.GetByUsername(context.Background(), username)
	if err != nil {
		http.Error(w, "user array creation error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, user)

}

// DeleteByUsername deletes a user based on their username.
// @Summary Delete a user by username
// @Description Delete a user by their username
// @Tags users
// @Produce json
// @Param username path string true "Username of the user to delete"
// @Success 200 {string} string "Users deleted"
// @Failure 400 {string} string "empty name error"
// @Failure 417 {string} string "user deletion error"
// @Router /users/{username} [delete]
func (uc UserCntrl) UpdateByUsername(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	pathSpl := strings.Split(r.URL.Path, "/")
	username := pathSpl[len(pathSpl)-1]

	if username == "" {
		http.Error(w, "empty name error", http.StatusBadRequest)
		return
	}

	err = uc.us.UpdateByUsername(context.Background(), username, &user)
	if username == "" {
		http.Error(w, "user update error"+err.Error(), http.StatusExpectationFailed)
		return
	}

	fmt.Fprint(w, "Users updated")

}

// Login authenticates a user and returns a JWT token.
// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 200 {string} string "User login succeed"
// @Failure 400 {string} string "request parsing error"
// @Failure 404 {string} string "login error"
// @Router /users/login [post]
func (uc UserCntrl) DeleteByUsername(w http.ResponseWriter, r *http.Request) {
	pathSpl := strings.Split(r.URL.Path, "/")
	username := pathSpl[len(pathSpl)-1]

	if username == "" {
		http.Error(w, "empty name error", http.StatusBadRequest)
		return
	}

	err := uc.us.DeleteByUsername(context.Background(), username)
	if username == "" {
		http.Error(w, "user deletion error"+err.Error(), http.StatusExpectationFailed)
		return
	}

	fmt.Fprint(w, "Users deleted")

}

// Login authenticates a user and returns a JWT token.
// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 200 {string} string "User login succeed"
// @Failure 400 {string} string "request parsing error"
// @Failure 404 {string} string "login error"
// @Router /users/login [post]
func (uc UserCntrl) Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}
	if user.Username == "" {
		http.Error(w, "empty name error", http.StatusBadRequest)
		return
	}

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	_, token, err := tokenAuth.Encode(map[string]interface{}{"username": user.Username})
	if err != nil {
		http.Error(w, "authentication internal error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	err = uc.us.Login(context.Background(), user.Username, token)
	if err != nil {
		http.Error(w, "login error:"+err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "User login succeed")

}

// Logout invalidates a user's JWT token.
// @Summary Logout a user
// @Description Invalidate a user's JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param token body string true "JWT token to invalidate"
// @Success 200 {string} string "user logout succeed"
// @Failure 400 {string} string "bad token error"
// @Router /users/logout [post]
func (uc UserCntrl) Logout(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	token := user.Token
	if token == "" {
		http.Error(w, "bad token error:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = uc.us.Logout(context.Background(), token)
	if token == "" {
		http.Error(w, "user logout error:"+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "user logout succeed")
}

// CreateUser creates a new user with the provided details.
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User to create"
// @Success 200 {string} string "user created successfully"
// @Failure 400 {string} string "request parsing error"
// @Failure 500 {string} string "user creation error"
// @Router /users [post]
func (uc UserCntrl) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = uc.us.CreateUser(context.Background(), &user)
	if err != nil {
		http.Error(w, "user creation error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "user created successfully")

}
