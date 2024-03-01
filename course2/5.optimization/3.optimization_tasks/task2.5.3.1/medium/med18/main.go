package main

import "fmt"

func main() {
	groupSizes := []int{3, 3, 3, 3, 3, 1, 3}
	fmt.Println(groupThePeople(groupSizes))
}

func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int)
	var res [][]int
	for i, groupN := range groupSizes {
		m[groupN] = append(m[groupN], i)
	}
	for groupN, group := range m {
		q := len(group) / groupN
		for i := 0; i < q; i++ {
			res = append(res, group[i*groupN:(i*groupN)+groupN])
		}
	}
	return res
}
