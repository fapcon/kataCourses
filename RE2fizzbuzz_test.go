package main

import "testing"

type testDataa struct {
	res int
	exp string
}

func TestFizzBuzz(t *testing.T) {
	testCases := []testDataa {
		{5, "buzz"},
		{3, "fizz"},
		{15, "fizz buzz"},
	}
	for _, tc := range testCases {
		result := fizzbuzz(tc.res)
		if result != tc.exp {
			t.Errorf("Unexpected result. Expected: %d, Got: %d", tc.exp, result)
		}
	}
}
