package main

import (
	"fmt"
	"math"
)

var a, result float64
func main() {
	result := Sqrt(4)
	fmt.Println(result)
}

func Sqrt(x float64) float64 {
	result = math.Sqrt(x)
	return result
}
