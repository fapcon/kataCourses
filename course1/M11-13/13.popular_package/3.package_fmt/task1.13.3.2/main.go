package main

import "fmt"

func main() {
	var num int = 10
	var str string = "Hello"

	fmt.Println(getVariableType(num)) // Вывод: "int"
	fmt.Println(getVariableType(str)) // Вывод: "string"
}

// Функция для получения типа переменной
func getVariableType(variable interface{}) string {
	var res string
	switch variable.(type) {
	case int:
		res = "int"
	case string:
		res = "string"
	case float64:
		res = "float64"
	case bool:
		res = "bool"
	case rune:
		res = "rune"
	}
	res = fmt.Sprintf("%s", res)
	return res
}
