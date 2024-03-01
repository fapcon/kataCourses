package main

import (
	"fmt"
)

func main() {
	q := []int{7, 5, 5, 8, 3}
	fmt.Println(processQueries(q, 8))
}

func processQueries(queries []int, m int) []int {
	var P, res []int
	var pos int
	for i := 0; i < m; i++ {
		P = append(P, i+1)
	}
	for i := 0; i < len(queries); i++ {
		for j := 0; j < len(P); j++ {
			if P[j] == queries[i] {
				res = append(res, j)
				pos = j
				//r = P[j]
				element := P[pos]
				P = append([]int{element}, append(P[:pos], P[pos+1:]...)...)
			}
		}
	}
	return res
}
