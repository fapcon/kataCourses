package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID       int    `db:"id" db_ops:"create"`
	Username string `db:"username" db_ops:"create,update"`
	Email    string `db:"email" db_ops:"create,update"`
	Address  string `db:"address" db_ops:"update"`
	Status   int    `db:"status" db_ops:"create,update"`
	Delete   string `db:"delete" db_ops:"delete"`
}

func SimpleGetFieldsPointers(u interface{}) []interface{} {
	v := reflect.ValueOf(u).Elem()
	res := []interface{}{}
	for i := 0; i < v.NumField(); i++ {
		fieldT := v.Type().Field(i)
		tag := fieldT.Tag.Get("db")
		if tag == "id" || tag == "delete" {
			continue
		}
		f := v.Field(i).Addr().Interface()
		res = append(res, f)
	}
	return res
}

func main() {
	user := User{
		ID:       1,
		Username: "JohnDoe",
		Email:    "johndoe@example.com",
		Address:  "123 Main St",
		Status:   1,
		Delete:   "yes",
	}

	pointers := SimpleGetFieldsPointers(&user)
	fmt.Println(pointers)
}
