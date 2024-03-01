package main

import (
	"math"
	"strconv"
)

func main() {

}

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	x, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0, err
	}
	y, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, err
	}
    if x == 0 {
		return 0, err
	}
	a := math.Round(((y/x*100-100)*100)/100)
	return a, err
}