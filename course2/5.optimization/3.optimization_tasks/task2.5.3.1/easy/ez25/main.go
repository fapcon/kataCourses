package main

import "fmt"

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}

func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		res = append(res[:index[i]], append([]int{nums[i]}, res[index[i]:]...)...)
	}
	return res
}
