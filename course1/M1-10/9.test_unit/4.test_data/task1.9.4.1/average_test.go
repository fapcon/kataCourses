package main

import (
	"math/rand"
	"testing"
)
var expected float64

func TestAverage(t *testing.T) {
testData := generateSlice(10)
result := average(testData)
var sum float64
for i:=0; i<len(testData);i++ {
	sum = sum + testData[i]
}
expected = sum/float64(len(testData))
if result != expected {
	t.Errorf("Ожидалось: %g, получено: %g", expected, result)
}
}

func generateSlice(size int) []float64 {
	var td []float64
	for i:=0;i<size;i++ {
		td=append(td, rand.Float64())
	}
	return td
}
