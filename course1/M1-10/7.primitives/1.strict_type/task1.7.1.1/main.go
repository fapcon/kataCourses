package main

import "fmt"

func main() {
	var name, city string
	var age int
	fmt.Println("Введите ваше имя: ")
	fmt.Scanln(&name)
	fmt.Println("Введите ваш возраст: ")
	fmt.Scanln(&age)
	fmt.Println("Введите ваш город:")
	fmt.Scanln(&city)
	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Город:", city)
}
