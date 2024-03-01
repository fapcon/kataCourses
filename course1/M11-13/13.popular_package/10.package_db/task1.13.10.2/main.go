package main

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {

}

// Создание таблицы пользователей
func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		err.Error()
	}
	defer db.Close()

	_, err = db.Exec(("CREATE TABLE users (id INTEGER PRIMARY KEY, username TEXT, email TEXT)"))
	if err != nil {
		err.Error()
	}
	return err
}

// Выборка пользователя из таблицы
func SelectUser(userID int) (User, error) {
	var user User
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return user, err
	}
	defer db.Close()

	query, args, err := PrepareQuery("select", "users", User{ID: userID})
	if err != nil {
		return user, err
	}

	rows := db.QueryRow(query, args...)
	err = rows.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		err.Error()
	}
	return user, err
}

// Обновление информации о пользователе
func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("update", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Удаление пользователя из таблицы
func DeleteUser(userID int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("delete", "users", User{ID: userID})
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Функция для подготовки запроса
func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	switch strings.ToLower(operation) {
	case "select":
		query, args, _ := sq.Select("*").From(table).Where(sq.Eq{"id": user.ID}).ToSql()
		return query, args, nil
	case "insert":
		query, args, _ := sq.Insert(table).Columns("username", "email").Values(user.Username, user.Email).ToSql()
		return query, args, nil
	case "update":
		query, args, _ := sq.Update(table).Set("username", user.Username).Set("email", user.Email).Where(sq.Eq{"id": user.ID}).ToSql()
		return query, args, nil
	case "delete":
		query, args, _ := sq.Delete(table).Where(sq.Eq{"id": user.ID}).ToSql()
		return query, args, nil
	}
	return "", nil, nil
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("insert", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
