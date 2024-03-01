package main

import "fmt"

func main() {

}

func PrintNumbers(numbers ...int) {
	for _, numbers := range numbers {
		fmt.Println(numbers)
	}
}
