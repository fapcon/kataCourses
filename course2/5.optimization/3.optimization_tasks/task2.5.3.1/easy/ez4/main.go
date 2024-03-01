package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
}

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}
