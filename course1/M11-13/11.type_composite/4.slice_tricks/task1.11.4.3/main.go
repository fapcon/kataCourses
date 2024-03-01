package main

import "fmt"

func main() {

}

func RemoveExtraMemory(xs []int) []int {
	var res []int
	if cap(xs) > len(xs) {
		res := make([]int, len(xs), len(xs))
		copy(xs, res)
		fmt.Println(cap(res))
	}

	return res
}
