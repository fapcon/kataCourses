package main

import "fmt"

func multiplier(factor float64) func(float64) float64 {
	var mul float64
	mul = 1
	return func(x float64) float64 {
		mul = x*factor
		return mul
	}
}

func main() {
	// Пример использования функции multiplier
	m := multiplier(2.5)
	result := m(10)
	fmt.Println(result) // Вывод: 25
}

