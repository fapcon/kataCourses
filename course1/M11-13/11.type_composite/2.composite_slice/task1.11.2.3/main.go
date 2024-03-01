package main

import "fmt"

func bitwiseXOR(n, res int) int {
	x := n ^ res
	return x
}

func findSingleNumber(numbers []int) int {

	var res = numbers[0]

	for i := 1; i < len(numbers); i++ {
		res = bitwiseXOR(res, numbers[i])
	}
	return res
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	singleNumber := findSingleNumber(numbers)
	fmt.Println(singleNumber) // 5
}
