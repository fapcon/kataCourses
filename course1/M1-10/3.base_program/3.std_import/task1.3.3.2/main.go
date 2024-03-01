package main

import (
	"fmt"
	"math"
)

var a, res1 float64
func main() {
	res1 := Floor(4.12323463246)

	fmt.Println(res1)

}

func Floor(x float64) float64 {
	a = math.Floor(x)
	return a
}
