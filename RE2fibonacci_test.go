package main

import (
	"testing"
	"time"
)

type testData struct {
	res int
	exp int
}

func TestFibonacci(t *testing.T) {
	testCases := []testData {
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}
	for _, tc := range testCases {
		result := FibonacciRe(tc.res)
		if result != tc.exp {
			t.Errorf("Unexpected result. Expected: %d, Got: %d", tc.exp, result)
		}
	}
	for _, tc := range testCases {
		result := FibonacciBinet(tc.res)
		if result != tc.exp {
			t.Errorf("Unexpected result. Expected: %d, Got: %d", tc.exp, result)
		}
	}
}

func BenchmarkFibonacciRe(b *testing.B) {
	a := time.Now()
	for i := 0; i < b.N; i++ { // запускаем бенчмарк
		FibonacciRe(5)
	}
	c := time.Since(a)
	b.Errorf("%v", c)
}
func BenchmarkFibonacciBinet(b *testing.B) {
	a := time.Now()
	for i := 0; i < b.N; i++ { // запускаем бенчмар
		FibonacciBinet(10)
	}
	c := time.Since(a)
	b.Errorf("%v", c)
}