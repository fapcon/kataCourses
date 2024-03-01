package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
}

func shuffle(nums []int, n int) []int {
	var res []int
	s1 := nums[:n]
	s2 := nums[n:]
	for i := 0; i < n; i++ {
		res = append(res, s1[i])
		res = append(res, s2[i])
	}
	return res
}
