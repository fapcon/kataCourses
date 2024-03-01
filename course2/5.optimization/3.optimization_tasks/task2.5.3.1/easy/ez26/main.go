package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}

func decompressRLElist(nums []int) []int {
	var ans []int
	freq := make([]int, len(nums)/2)
	val := make([]int, len(nums)/2)
	res := make([][]int, len(freq))
	var i int
	var j int
	for i < len(nums) {
		freq[j] = nums[i]
		val[j] = nums[i+1]
		i += 2
		j++
	}
	for i := 0; i < len(freq); i++ {
		for j := 0; j < freq[i]; j++ {
			res[i] = append(res[i], val[i])
		}
	}
	for i := 0; i < len(res); i++ {
		ans = append(ans, res[i]...)
	}
	return ans
}
