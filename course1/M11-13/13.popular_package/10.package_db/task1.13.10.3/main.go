package main

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
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
	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
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

	err = PrepareQuery("select", "users", User{ID: userID}).(*sq.SelectBuilder).RunWith(db).QueryRow().Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return user, err
	}

	return user, err
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = PrepareQuery("insert", "users", user).(*sq.InsertBuilder).RunWith(db).Exec()
	if err != nil {
		return err
	}
	return err
}

// Обновление информации о пользователе
func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	_, err = PrepareQuery("update", "users", user).(*sq.UpdateBuilder).RunWith(db).Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("User updated successfully")
	return err
}

// Удаление пользователя из таблицы
func DeleteUser(userID int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = PrepareQuery("delete", "users", User{ID: userID}).(*sq.DeleteBuilder).RunWith(db).Exec()
	if err != nil {
		return err
	}
	return err
}

// Функция для подготовки запроса
func PrepareQuery(operation string, table string, user User) interface{} {

	switch strings.ToLower(operation) {
	case "select":
		res := sq.Select("*").From(table).Where(sq.Eq{"id": user.ID})
		return &res
	case "insert":
		res := sq.Insert(table).Columns("name", "age").Values(user.Name, user.Age)
		return &res
	case "update":
		res := sq.Update(table).Set("name", user.Name).Set("age", user.Age).Where(sq.Eq{"id": user.ID})
		return &res
	case "delete":
		res := sq.Delete(table).Where(sq.Eq{"id": user.ID})
		return &res
	}
	return nil
}
