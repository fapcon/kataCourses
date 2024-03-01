package main

import "fmt"

func main() {
	fmt.Println(smallestEvenMultiple(77))
}

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	} else {
		return n * 2
	}
}
