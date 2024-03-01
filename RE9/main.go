package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

//репозиторий

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func CreateTable(db *sql.DB) error {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS USERS (EMAIL TEXT PRIMARY KEY, PASSWORD TEXT, NAME TEXT, AGE INTEGER);`)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) CreateUser(user User) error {
	var count int
	err := ur.db.QueryRow(`SELECT COUNT(*) FROM USERS WHERE EMAIL = ? OR (AGE < 18 AND AGE = ?)`, user.Email, user.Age).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("user already exists")
	}

	_, err = ur.db.Exec(`INSERT INTO USERS (EMAIL, PASSWORD, NAME, AGE) VALUES (?, ?, ?, ?)`, user.Email, user.Password, user.Name, user.Age)

	return err
}

func (ur *UserRepository) GetAllUsers() ([]User, error) {
	rows, err := ur.db.Query(`SELECT EMAIL, PASSWORD, NAME, AGE FROM USERS`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Email, &user.Password, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

type CachedDatabase struct {
	DB    *UserRepository
	mu    sync.RWMutex
	cache map[string]interface{}
}

func NewCacheDB(db *UserRepository) *CachedDatabase {
	return &CachedDatabase{
		DB:    db,
		cache: make(map[string]interface{}),
	}
}

func (cd *CachedDatabase) Get(key string) (interface{}, bool) {
	cd.mu.RLock()
	defer cd.mu.RUnlock()
	value, ok := cd.cache[key]
	return value, ok
}

func (cd *CachedDatabase) Set(key string, value interface{}) {
	cd.mu.Lock()
	defer cd.mu.Unlock()
	cd.cache[key] = value
}

func (cd *CachedDatabase) GetCacheAllUsers() ([]User, error) {
	if value, ok := cd.Get("userss"); ok {
		if users, ok := value.([]User); ok {
			fmt.Println("cache used")
			return users, nil
		}
	}

	users, err := cd.DB.GetAllUsers()
	if err != nil {
		return nil, err
	}

	cd.Set("userss", users)

	return users, nil
}

//сервис

type UserService struct {
	UserRepository *UserRepository
}

func NewUserService(db *UserRepository) *UserService {
	return &UserService{db}
}

func (us *UserService) CreateUser(user User) error {
	return us.UserRepository.CreateUser(user)
}

func (us *UserService) GetAllUsers() ([]User, error) {
	return us.UserRepository.GetAllUsers()
}

//контроллеры

type UserHandler struct {
	UserService *UserService
	CacheDB     *CachedDatabase
}

func NewUserHandler(userService *UserService, cache *CachedDatabase) *UserHandler {
	return &UserHandler{userService, cache}
}

func (uh *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Email:    "@@@@@@@@@@",
		Password: "******!!!!!",
		Name:     "qwerty",
		Age:      73,
	}

	if err := uh.UserService.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, _ := uh.CacheDB.GetCacheAllUsers()
	users = append(users, user)
	uh.CacheDB.Set("userss", users)

	w.Write([]byte("user successfully registered"))

	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uh.CacheDB.GetCacheAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	router := chi.NewRouter()

	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}

	rep := NewUserRepository(db)

	service := NewUserService(rep)

	cache := NewCacheDB(rep)

	handler := NewUserHandler(service, cache)

	err = CreateTable(db)

	if err != nil {
		log.Fatal(err)
	}

	router.Get("/register", handler.RegisterUserHandler)

	router.Get("/getallusers", handler.GetAllUsersHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
