package main

import "fmt"

func main() {
	fmt.Println(decode([]int{6, 2, 7, 3}, 4))
}

func decode(encoded []int, first int) []int {
	for i, v := range encoded {
		encoded[i] = first
		first = first ^ v
	}

	encoded = append(encoded, first)

	return encoded
}
