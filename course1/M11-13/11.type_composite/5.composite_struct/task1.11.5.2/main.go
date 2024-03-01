package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"strconv"
)

func main() {
	users := getAnimals()
	result := preparePrint(users)
	fmt.Println(result)
}

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	list := make([]Animal, 3)
	for i := 0; i < len(list); i++ {
		gofakeit.Struct(&list[i])
	}
	return list
}

func preparePrint(animals []Animal) string {
	var res string
	st := make([]string, 3)
	for i := 0; i < len(animals); i++ {
		st[i] = "Тип: " + animals[i].Type + ", Имя: " + animals[i].Name + ", Возраст: " + strconv.Itoa(animals[i].Age)
	}
	for i := 0; i < len(st); i++ {
		res = res + st[i] + "\n"
	}
	return res
}
