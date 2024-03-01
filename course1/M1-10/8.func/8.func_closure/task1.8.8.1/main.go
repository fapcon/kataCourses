package main

import "fmt"

func createCounter() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func main() {
	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}