package main

import "fmt"

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}

func numIdenticalPairs(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {

		}
	}
	return res
}
