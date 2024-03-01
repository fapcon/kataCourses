package main

import "fmt"

func main() {
	fmt.Println(InsertAfterIDX([]int{1, 2, 3, 4, 5, 6, 7}, 2, 111, 112))
}

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	var a, b []int
	b = append(b, x...)
	a = append(xs[:idx+1], append(b, xs[idx+1:]...)...)
	return a
}
