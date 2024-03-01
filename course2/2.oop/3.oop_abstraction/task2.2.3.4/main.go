package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type SQLiteGenerator struct {
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
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

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLiteGenerator
}

type Tabler interface {
	TableName() string
}

func (u User) TableName() string {
	return "users"
}

func NewMigrator(db *sql.DB, generator SQLiteGenerator) *Migrator {
	migrator := &Migrator{db: db, sqlGenerator: generator}
	return migrator
}

func (m *Migrator) Migrate(model ...Tabler) error {
	for _, table := range model {
		migrate := m.sqlGenerator.CreateTableSQL(table)
		_, err := m.db.Exec(migrate)
		if err != nil {
			return fmt.Errorf("failed migrate %s: %v", table.TableName(), err)
		}
	}
	return nil
}

// Основная функция
func main() {
	// Подключение к SQLite БД
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Создание мигратора с использованием вашего SQLGenerator
	YourSQLGeneratorInstance := SQLiteGenerator{}
	migrator := NewMigrator(db, YourSQLGeneratorInstance)

	// Миграция таблицы User
	if err := migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
