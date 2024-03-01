package main

import "fmt"

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1, 3, 3}))
}

func getConcatenation(nums []int) []int {
	ans := append(nums, nums...)
	return ans
}
