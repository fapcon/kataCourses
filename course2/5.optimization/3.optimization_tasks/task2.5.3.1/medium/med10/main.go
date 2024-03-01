package main

import (
	"fmt"
	"sort"
)

func main() {
	titles := "AAB"
	fmt.Println(numTilePossibilities(titles))
}

func numTilePossibilities(tiles string) int {
	res := 0
	arr := []byte(tiles)
	bites := make([]bool, len(tiles))
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var r func()
	r = func() {
		for i := 0; i < len(arr); i++ {
			if bites[i] {
				continue
			}
			bites[i] = true
			res++
			r()
			bites[i] = false
			for ; i+1 < len(arr) && arr[i] == arr[i+1]; i++ {
			}
		}
	}
	r()
	return res
}
