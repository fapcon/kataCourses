package main

import (
	"fmt"
	"math"
)

func main() {
	points := [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}
	queries := [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}
	fmt.Println(countPoints(points, queries))
}

func countPoints(points [][]int, queries [][]int) []int {
	var res []int
	for _, val := range queries {
		i := 0
		a, b, r := float64(val[0]), float64((val[1])), float64(val[2])
		for _, d := range points {
			distance := math.Sqrt(math.Pow(float64(d[0])-a, 2) + math.Pow(float64(d[1])-b, 2))
			if distance <= r {
				i++
			}
		}
		res = append(res, i)
	}
	return res
}
