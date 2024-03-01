package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var m int
	for i, e := range candies {
		if i == 0 || e > m {
			m = e
		}
	}
	var res []bool
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= m {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
