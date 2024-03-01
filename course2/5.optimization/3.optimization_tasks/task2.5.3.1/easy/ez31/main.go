package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortHeights([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
}

func sortHeights(names []string, heights []int) []string {
	//var res []string
	//m := make(map[string]int, len(names))
	//for i := 0; i < len(names); i++ {
	//	m[names[i]] = heights[i]
	//}
	//sort.Ints(heights)

	for i := 0; i < len(names)-1; i++ {
		for j := 0; j < len(names)-i-1; j++ {
			if heights[j] < heights[j+1] {
				heights[j], heights[j+1] = heights[j+1], heights[j]
				names[j], names[j+1] = names[j+1], names[j]
			}
		}
	}
	fmt.Println(heights)
	return names
}
