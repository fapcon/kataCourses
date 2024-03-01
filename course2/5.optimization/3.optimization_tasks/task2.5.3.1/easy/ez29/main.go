package main

import "fmt"

func main() {
	fmt.Println(xorOperation(4, 3))
}

func xorOperation(n int, start int) int {
	nums := make([]int, n)
	nums[0] = start
	res := 0
	for i := 1; i < len(nums); i++ {
		nums[i] = start + 2*i
	}
	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
