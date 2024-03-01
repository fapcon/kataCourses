package main

import "testing"

type testData struct {
	a   int
	b   int
	exp int
}

type testData2 struct {
	arr []int
	exp int
}

func TestBitwiseXOR(t *testing.T) {
	testCases := []testData{
		{0, 0, 0},
		{0, 1, 1},
		{1, 1, 0},
	}
	for _, tc := range testCases {
		result := bitwiseXOR(tc.a, tc.b)
		if result != tc.exp {
			t.Errorf("Unexpected result. Expected: %d, Got: %d", tc.exp, result)
		}
	}
}

func TestFindSingleNumber(t *testing.T) {
	testCases := []testData2{
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1}, 5},
		{[]int{5, 4, 3, 2, 1, 2, 3, 4, 5}, 1},
	}
	for _, tc := range testCases {
		result := findSingleNumber(tc.arr)
		if result != tc.exp {
			t.Errorf("Unexpected result. Expected: %d, Got: %d", tc.exp, result)
		}
	}
}
