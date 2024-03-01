package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}

func runningSum(nums []int) []int {
	var sum int
	var res []int
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		res = append(res, sum)
	}
	return res
}
