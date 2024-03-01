package main

import "fmt"

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}

func findKthPositive(arr []int, k int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid]-mid-1 < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left + k
}
