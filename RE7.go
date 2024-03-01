package main

import (
	"crypto/tls"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type User struct {
	ID        int64   `json:"id"`
	Email     string  `json:"email"`
	Amount    int     `json:"amount"`
	Profile   Profile `json:"profile"`
	Password  string  `json:"-"`
	Username  string  `json:"username"`
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
}
type Profile struct {
	Dob        string `json:"dob"`
	Avatar     string `json:"avatar"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	StaticData string `json:"-"`
}
type UserResponse struct {
	Records []*User `json:"records"`
}
type UserPublic struct {
	ID        int64         `json:"id"`
	Email     *string       `json:"email,omitempty"`
	Amount    int           `json:"amount"`
	Profile   ProfilePublic `json:"profile"`
	CreatedAt string        `json:"createdAt"`
	CreatedBy string        `json:"createdBy"`
}
type ProfilePublic struct {
	Dob       *string `json:"dob,omitempty"`
	Avatar    *string `json:"avatar,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/users", UserHandler)
	http.ListenAndServe(":8080", r)
}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	cl := http.Client{Transport: tr}
	resp, err := cl.Get("https://demo.apistubs.io/api/v1/users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var users UserResponse
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users.Records[0].Amount = 50001
	var publicUsers []UserPublic
	for _, user := range users.Records {
		publicUser := UserPublic{
			ID:        user.ID,
			Amount:    user.Amount,
			CreatedAt: user.CreatedAt,
			CreatedBy: user.CreatedBy,
		}
		if user.Amount <= 50000 {
			publicUser.Email = &user.Email
			publicUser.Profile.FirstName = &user.Profile.FirstName
			publicUser.Profile.LastName = &user.Profile.LastName
			publicUser.Profile.Avatar = &user.Profile.Avatar
		}
		publicUsers = append(publicUsers, publicUser)
	}
	data, err := json.Marshal(publicUsers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
