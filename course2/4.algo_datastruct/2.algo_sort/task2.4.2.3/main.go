package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func generateUsers(n int) []User {
	gofakeit.Seed(time.Now().UnixNano())
	users := make([]User, n)
	for i := range users {
		users[i] = User{
			ID:   gofakeit.Number(1, 100),
			Name: gofakeit.FirstName(),
			Age:  gofakeit.Number(1, 80),
		}
	}
	return users
}

// Функция слияния двух отсортированных массивов пользователей
func Merge(arr1 []User, arr2 []User) []User {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}
	arrR := append(arr1, arr2...)
	n := len(arrR)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arrR[j].ID < arrR[minIndex].ID {
				minIndex = j
			}
		}
		arrR[i], arrR[minIndex] = arrR[minIndex], arrR[i]
	}
	return arrR
}

func main() {
	users1 := generateUsers(10)
	users2 := generateUsers(0)
	fmt.Println(Merge(users1, users2))
}
