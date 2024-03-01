package main

import "math"


func main() {

}
var CalculateCircleArea = func(radius float64) float64 {
	return math.Pi*radius*radius
}
var CalculateRectangleArea = func(width, height float64) float64 {
	return width*height
}
var CalculateTriangleArea = func(base, height float64) float64 {
	return 0.5*base*height
}




