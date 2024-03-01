package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	fmt.Println(checkArithmeticSubarrays(nums, l, r))
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var res []bool
	for a := 0; a < len(l); a++ {
		var sub []int
		start := l[a]
		end := r[a]
		sub = append(sub, nums[start:end+1]...)
		//fmt.Println(sub)
		//sort.Ints(sub)
		//fmt.Println(sub)
		res = append(res, trueSec(sub))
	}
	return res
}

func trueSec(a []int) bool {
	for i := 1; i < len(a)-1; i++ {
		if a[i+1]-a[i] != a[i]-a[i-1] {
			return false
		}
	}
	return true
}
