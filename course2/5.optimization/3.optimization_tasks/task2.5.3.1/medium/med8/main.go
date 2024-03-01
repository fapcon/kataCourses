package main

import "fmt"

func main() {
	edges := [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}
	fmt.Println(findSmallestSetOfVertices(6, edges))
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	arr := make([]int, n)
	for _, edge := range edges {
		arr[edge[1]] = 1
	}
	var res []int
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
