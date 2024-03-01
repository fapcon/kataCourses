package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		err.Error()
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, name VARCHAR(255) NOT NULL, age INT);")
	if err != nil {
		err.Error()
	}
	return err
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		err.Error()
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO  user  (name,age) VALUES (?,?)", user.Name, user.Age)
	if err != nil {
		err.Error()
	}
	return err
}

func SelectUser(id int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		err.Error()
	}
	defer db.Close()
	rows, err := db.Query("SELECT id , name , age FROM  user  WHERE id=?", id)
	if err != nil {
		err.Error()
	}
	//defer rows.Close()
	var res User
	for rows.Next() {
		err = rows.Scan(&res.ID, &res.Name, &res.Age)
		if err != nil {
			err.Error()
		}
	}
	return res, err
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		err.Error()
	}
	defer db.Close()

	_, err = db.Exec("UPDATE user SET id=?, name=?, age=? WHERE id=?", user.ID, user.Name, user.Age, user.ID)
	if err != nil {
		err.Error()
	}

	return err
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		err.Error()
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	return err
}
