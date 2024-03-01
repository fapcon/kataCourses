package main

import "testing"

type testData struct {
	res int
	exp int
}

func TestFibonacci(t *testing.T) {
testCases := []testData{
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
}
for _, tc := range testCases {
	result := Fibonacci(tc.res)
	if result != tc.exp {
		t.Errorf("Unexpected result. Expected: %d, Got: %d", tc.exp, result)
	}
}
}

