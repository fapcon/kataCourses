package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(differenceOfSum([]int{1, 15, 6, 3}))
}

func differenceOfSum(nums []int) int {
	var elmSum, digSum int
	var str []string
	for i := 0; i < len(nums); i++ {
		elmSum = elmSum + nums[i]
		str = append(str, strings.Split(strconv.Itoa(nums[i]), "")...)
	}
	sti := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		sti[i], _ = strconv.Atoi(str[i])
		digSum = digSum + sti[i]
	}
	if elmSum > digSum {
		return elmSum - digSum
	} else {
		return digSum - elmSum
	}
}
