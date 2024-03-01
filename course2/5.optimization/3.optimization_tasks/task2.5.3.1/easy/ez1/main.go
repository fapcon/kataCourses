package main

import "fmt"

func main() {
	fmt.Println(tribonacci(4))
}

func tribonacci(n int) int {
	array := make([]int, n+1)
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	array[0] = 0
	array[1] = 1
	array[2] = 1
	var res int
	for i := 3; i <= n; i++ {
		res = array[i-1] + array[i-2] + array[i-3]
		array[i] = res
	}
	return res
}
