package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSum(5678))
}

func minimumSum(num int) int {
	digits := []int{}

	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	sort.Ints(digits)

	num1 := digits[0]*10 + digits[2]
	num2 := digits[1]*10 + digits[3]

	return num1 + num2
}
