package main

import (
	"fmt"
	"github.com/fapcon/mymath"
)

func Abs(x float64) float64 {
	return mymath.Abs(x)
}
func Yn(x int, y float64) float64 {
	return mymath.Yn(x, y)
}
func Max(x, y float64) float64 {
	return mymath.Max(x, y)
}
func Sqrt(x float64) float64 {
	return mymath.Sqrt(x)
}

func main() {
	x := 4.0
	y := 2.0

	Abs := mymath.Abs(x)
	Max := mymath.Max(x, y)
	Sqrt := mymath.Sqrt(x)
	Yn := mymath.Yn(int(x), y)
	fmt.Println(Abs)
	fmt.Println(Max)
	fmt.Println(Sqrt)
	fmt.Println(Yn)
}
