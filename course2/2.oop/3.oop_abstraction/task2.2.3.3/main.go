package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"reflect"
	"strings"
)

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type Tabler interface {
	TableName() string
}

type SQLiteGenerator struct {
}

type GoFakeitGenerator struct {
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	var fields []string
	userType := reflect.TypeOf(User{})
	for i := 0; i < userType.NumField(); i++ {
		//res = append(res, reflect.TypeOf(e).Field(i).Tag)
		field := userType.Field(i)
		dbField := field.Tag.Get("db_field")
		dbType := field.Tag.Get("db_type")
		fields = append(fields, fmt.Sprintf("%s %s", dbField, dbType))
	}
	return "CREATE TABLE IF NOT EXISTS " + table.TableName() + " (" + strings.Join(fields, ", ") + ");"
}

func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	var fields []string
	userType := reflect.TypeOf(model).Elem()
	userVal := reflect.ValueOf(model).Elem()
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		dbField := field.Tag.Get("db_field")
		fields = append(fields, fmt.Sprintf("%s", dbField))
	}
	return fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v, '%v', '%v', '%v')", model.TableName(), strings.Join(fields, ", "), userVal.Field(0), userVal.Field(1), userVal.Field(2), userVal.Field(3))
}

func (u User) TableName() string {
	return "users"
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	return User{ID: gofakeit.IntRange(1, 10000), FirstName: gofakeit.FirstName(), LastName: gofakeit.LastName(), Email: gofakeit.Email()}
}

func GenerateUserInserts(a int) []string {
	res := make([]string, a)
	users := make([]User, a)
	for i := 0; i < a; i++ {
		users[i] = User{ID: gofakeit.IntRange(1, 10000), FirstName: gofakeit.FirstName(), LastName: gofakeit.LastName(), Email: gofakeit.Email()}
	}
	var fields []string
	userType := reflect.TypeOf(users).Elem()

	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		dbField := field.Tag.Get("db_field")
		fields = append(fields, fmt.Sprintf("%s", dbField))
	}
	for i := 0; i < a; i++ {
		res = append(res, fmt.Sprintf("INSERT INTO users (%v) VALUES (%v, '%v', '%v', '%v')", strings.Join(fields, ", "), users[i].ID, users[i].FirstName, users[i].LastName, users[i].Email))
	}
	return res
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}

	queries := GenerateUserInserts(34)
	for _, query := range queries {
		fmt.Println(query)
	}
}
