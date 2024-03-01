package main

import (
	"math"
)

func main() {
}
func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	a := math.Round(((finalValue/initialValue)*100-100)*100)/100
	return a
}
