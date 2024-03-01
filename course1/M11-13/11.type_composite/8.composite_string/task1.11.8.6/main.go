package main

import "fmt"

func main() {
	count := CountVowels("Привет, мир!")
	fmt.Println(count) // Вывод: 3
}

var vow = []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', 'а', 'у', 'о', 'и', 'э', 'ы', 'я', 'ю', 'е', 'ё', 'А', 'Е', 'Ё', 'И', 'О', 'У', 'GEOapiSwagger', 'Э', 'Ю', 'Я'}

func CountVowels(str string) int {
	i := 0
	r := []rune(str)
	for _, val := range r {
		for j := 0; j < len(vow); j++ {
			if val == vow[j] {
				i++
			}
		}

	}
	return i
}
