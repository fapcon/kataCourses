package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := sum([8]int{1,2,3,4,5,6,7,8})
	xs := [8]int{1,2,3,4,5,6,7,8}
	sum := 0
	for _, xs := range xs {
		sum = sum+xs
}
	if result != sum {
	fmt.Errorf("expected: %d, got: %d", sum, result)
	}
}

func TestAverage(t *testing.T) {
	result := average([8]int{1,2,3,4,5,6,7,8})
	xs := [8]int{1,2,3,4,5,6,7,8}
	sum := 0
	for _, xs := range xs {
		sum = sum+xs
	}
	av := float64(sum)/8
	if result != av {
		fmt.Errorf("expected: %f, got: %f", av, result)
	}
}

func TestAverageFloat(t *testing.T) {
	result := averageFloat([8]float64{1.1,2.2,3.3,4.4,5.5,6.6,7.7,8.8})
	xs := [8]float64{1.1,2.2,3.3,4.4,5.5,6.6,7.7,8.8}
	sum := 0.0
	for _, xs := range xs {
		sum = sum+xs
	}
	av := sum/8
	if result != av {
		fmt.Errorf("expected: %f, got: %f", av, result)
	}
}

func TestReverse(t *testing.T) {
	result := reverse([8]int{1,2,3,4,5,6,7,8})
	xs := [8]int{1,2,3,4,5,6,7,8}
	var x [8]int
	temp := 0
	for i:=0; i<8; i++ {
		temp = xs[i]
		x[i]=xs[7-i]
		x[7-i]=temp
	}
	if result != x {
		fmt.Errorf("expected: %v, got: %v", x, result)
	}
}

