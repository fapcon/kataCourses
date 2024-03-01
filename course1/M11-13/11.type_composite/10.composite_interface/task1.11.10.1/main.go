package main

import "fmt"

func main() {
	var i interface{} = 42
	fmt.Println(getType(i)) // Вывод: "int"

	var j interface{} = "Hello, World!"
	fmt.Println(getType(j)) // Вывод: "string"

	var k interface{} = []int{1, 2, 3}
	fmt.Println(getType(k)) // Вывод: "[]int"

	var l interface{} = interface{}(nil)
	fmt.Println(getType(l)) // Вывод: "Пустой интерфейс"
}

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case []int:
		return "[]int"
	case string:
		return "string"
	case nil:
		return "Пустой интерфейс"
	case bool:
		return "bool"
	case float64:
		return "float64"
	case rune:
		return "rune"
	}
	return ""
}
