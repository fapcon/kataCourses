package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSum([][]int{{1, 2, 3, 10},
		{4, 5, 6, 20},
		{7, 8, 9, 30},
		{0, 10, 20, 30}}))
}

func maxSum(grid [][]int) int {
	var sum []int
	var s = 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			s = grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			sum = append(sum, s)
		}
	}
	sort.Ints(sum)
	fmt.Println(sum)
	return sum[len(sum)-1]
}
