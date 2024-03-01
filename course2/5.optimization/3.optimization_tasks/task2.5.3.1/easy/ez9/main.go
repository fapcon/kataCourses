package main

import "fmt"

func main() {
	fmt.Println(finalValueAfterOperations([]string{"++X", "++X", "X ++"}))
}

func finalValueAfterOperations(operations []string) int {
	res := 0
	for i := range operations {
		switch operations[i] {
		case "--X":
			res--
		case "X --":
			res--
		case "++X":
			res++
		case "X ++":
			res++
		}
	}
	return res
}
