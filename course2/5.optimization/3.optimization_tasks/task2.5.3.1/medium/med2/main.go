package main

import "fmt"

func main() {
	fmt.Println(sortE([][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}, 2))
}

func sortE(marks [][]int, k int) [][]int {
	for i := 0; i < len(marks)-1; i++ {
		if marks[i][k] < marks[i+1][k] {
			marks[i], marks[i+1] = marks[i+1], marks[i]
		}

	}
	return marks
}
