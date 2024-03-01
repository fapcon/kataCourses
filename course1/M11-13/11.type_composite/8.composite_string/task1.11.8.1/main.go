package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	bytes := countBytes("Привет, мир!")
	fmt.Println(bytes) // Вывод: 21

	// Пример использования функции countSymbols
	symbols := countSymbols("Привет, мир!")
	fmt.Println(symbols) // Вывод: 12
}

func countBytes(s string) int {
	b := []byte(s)
	return len(b)
}

func countSymbols(s string) int {
	return utf8.RuneCountInString(s)
}
