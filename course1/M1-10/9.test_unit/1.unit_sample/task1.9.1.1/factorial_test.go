package main

import "testing"

func TestFactorial(t *testing.T) {
	result1 := Factorial(5)
	if result1 != 120 {
		t.Errorf("Factorial(5) = %d; want 120", result1)
	}
	result2 := Factorial(0)
	if result2 != 1 {
		t.Errorf("Factorial(5) = %d; want 1", result2)
	}
	result3 := Factorial(1)
	if result3 != 1 {
		t.Errorf("Factorial(5) = %d; want 1", result3)
	}

}
