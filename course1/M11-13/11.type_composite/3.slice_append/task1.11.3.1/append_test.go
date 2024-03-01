package main

import (
	"fmt"
	"testing"
)

type testData struct {
	a   []int
	b   int
	exp []int
}

func TestAppendInt(t *testing.T) {
	testCases := []testData{
		{[]int{1, 2, 3}, 5, []int{1, 2, 3, 5}},
		{[]int{5, 4, 3, 2}, 1, []int{5, 4, 3, 2, 1}},
	}
	for _, tc := range testCases {
		result := appendInt(tc.a, tc.b)
		if len(result) != len(tc.exp) {
			fmt.Errorf("Unexpected result. Expected: %d, Got: %d", len(tc.exp), len(result))
		}
		for i, v := range result {
			if v != tc.exp[i] {
				fmt.Errorf("Unexpected result. Expected: %v, Got: %v", tc.exp, result)
			}
		}
	}
}
