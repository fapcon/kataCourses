package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type City struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	State string `db:"state"`
}

type CityRepository struct {
	db *sqlx.DB
}

func NewCityRepository(db *sqlx.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) Create(city *City) error {
	_, err := r.db.Exec(`INSERT INTO cities (id, name, state) VALUES (?, ?, ?)`, city.Id, city.Name, city.State)
	return err
}

func (r *CityRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM cities WHERE id = $1`, id)
	return err
}

func (r *CityRepository) Update(city *City) error {
	_, err := r.db.Exec(`UPDATE cities SET name = $1, state = $2 WHERE id = $3`, city.Name, city.State, city.Id)
	return err
}

func (r *CityRepository) List() ([]City, error) {
	rows, err := r.db.Query(`SELECT * FROM cities`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cities []City
	for rows.Next() {
		var city City
		err := rows.Scan(&city.Id, &city.Name, &city.State)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	//fmt.Println(cities)
	return cities, err
}

func CreateTables(db *sqlx.DB) error {
	_, err := db.Exec(`CREATE TABLE cities (id INTEGER NOT NULL PRIMARY KEY, name VARCHAR(30) NOT NULL, state VARCHAR(30) NOT NULL)`)
	if err != nil {
		log.Println(err)
	}
	return err
}

func generateCities() City {
	gofakeit.Seed(time.Now().UnixNano())
	city := City{
		Id:    gofakeit.Number(1, 100),
		Name:  gofakeit.City(),
		State: gofakeit.State(),
	}
	return city
}

func main() {
	db, err := sqlx.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = CreateTables(db)

	repo := NewCityRepository(db)

	city := generateCities()

	err = repo.Create(&city)

	fmt.Println("city created")

	if err != nil {
		log.Println(err)
	}

	list, err := repo.List()
	for _, n := range list {
		fmt.Println(n.Name)
	}
}
