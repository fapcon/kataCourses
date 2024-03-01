package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"strconv"
)

type User struct {
	Name string `fake:"{firstname}"`
	Age  int    `fake:"{number:18,60}"`
}

func main() {
	users := getUsers()
	result := preparePrint(users)
	fmt.Println(result)
}

func getUsers() []User {
	list := make([]User, 10)
	for i := 0; i < len(list); i++ {
		gofakeit.Struct(&list[i])
	}
	return list
}

func preparePrint(us []User) string {
	var res string
	st := make([]string, 10)
	for i := 0; i < len(us); i++ {
		st[i] = "Имя: " + us[i].Name + ", Возраст: " + strconv.Itoa(us[i].Age) + "\n"
	}
	for i := 0; i < len(st); i++ {
		res = res + st[i]
	}
	return res
}
