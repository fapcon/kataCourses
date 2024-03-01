package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findG([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3))
}

func findG(arr []int, a, b, c int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for k := 0; k < len(arr); k++ {
				if 0 <= i && i < j && j < k && k < len(arr) && math.Abs(float64(arr[i]-arr[j])) < float64(a) && math.Abs(float64(arr[j]-arr[k])) < float64(b) && math.Abs(float64(arr[i]-arr[k])) < float64(c) {
					res++
				}
			}
		}
	}
	return res
}
