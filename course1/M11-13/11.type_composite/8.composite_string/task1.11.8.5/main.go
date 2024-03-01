package main

import "fmt"

func main() {
	fmt.Println(ReverseString("Hello, world!")) // Вывод: "!dlrow ,olleH"
	fmt.Println(ReverseString("12345"))         // Вывод: "54321"
}

func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
