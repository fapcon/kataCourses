package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CompareRoundedValues(12.3456789, 98.7654321, 5))
}

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	var mul float64
	switch decimalPlaces {
	case 0: mul = 1
	case 1: mul = 10
	case 2: mul = 100
	case 3: mul = 1000
	case 4: mul = 10000
	case 5: mul = 100000
	case 6: mul = 1000000
	case 7: mul = 10000000
	case 8: mul = 100000000
	case 9: mul = 1000000000
	case 10: mul = 10000000000
	case 11: mul = 100000000000
	case 12: mul = 1000000000000
	case 13: mul = 10000000000000
	case 14: mul = 100000000000000
	case 15: mul = 1000000000000000
	case 16: mul = 10000000000000000
	case 17: mul = 100000000000000000
	case 18: mul = 1000000000000000000
	case 19: mul = 10000000000000000000



	}
	r1 := math.Round(a*mul)/mul
	r2 := math.Round(b*mul)/mul
	if r1 == r2 {
		isEqual = true
		difference = 0
	} else {
		isEqual = false
		difference = r1-r2
	}
	return isEqual, difference
}