package main

import "math"

func main() {

}

func hypotenuse(a, b float64) float64 {
	s := a*a + b*b
	res := math.Sqrt(s)
	return res
}
