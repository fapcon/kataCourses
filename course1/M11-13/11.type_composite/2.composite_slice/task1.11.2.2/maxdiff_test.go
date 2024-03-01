package main

import (
	"fmt"
	"testing"
)

func TestMaxDifference(t *testing.T) {
	numbers := []int{30,2,3,4,5,6,7,8,9,10,11,12,13,14}
	result := MaxDifference(numbers)
	exp := 0
	if len(numbers) == 0 || len(numbers) == 1 {
		exp = 0
	}
	max := numbers[0]
	min := numbers[0]
	for i:=0; i<len(numbers); i++ {
		if max<numbers[i] {
			max = numbers[i]
		}
		if min>numbers[i] {
			min = numbers[i]
		}
	}
	exp = max-min
	if result != exp {
		fmt.Errorf("expected: %v, got: %v", exp, result)
	}
}
