package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	fmt.Println(xorQueries(arr, queries))
}

func xorQueries(arr []int, queries [][]int) []int {
	var res []int
	var start, end int
	for i := 0; i < len(arr); i++ {
		start = queries[i][0]
		end = queries[i][1]
		//lenNums := end - start + 1
		var nums []int
		if i >= start && i <= end {
			for k := start; k < end+1; k++ {
				nums = append(nums, arr[k])
			}
		}
		fmt.Println(nums)
		x := nums[0]
		for j := 0; j < len(nums)-1; j++ {
			x = x ^ nums[j+1]
		}
		//x := arr[queries[i][0]] ^ arr[queries[i][1]]
		res = append(res, x)
	}
	return res
}
