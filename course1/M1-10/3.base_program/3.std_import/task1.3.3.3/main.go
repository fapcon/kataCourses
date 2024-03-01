package main

import (
	"fmt"
	"math"
)

var a, b, res1, res2 float64
func main() {
	res1 := Sin(4)
	res2 := Cos(4)
	fmt.Println(res1)
	fmt.Println(res2)
}

func Sin(x float64) float64 {
	a = math.Sin(x)
	return a
}

func Cos(x float64) float64 {
	b = math.Cos(x)
	return b
}

