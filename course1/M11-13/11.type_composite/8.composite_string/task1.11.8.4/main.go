package main

import "fmt"

func main() {
	result := concatStrings("")
	fmt.Println(result) // Вывод: "Hello world!"
}

func concatStrings(xs ...string) string {
	var res string
	for _, val := range xs {
		res = res + val
	}
	return res
}
