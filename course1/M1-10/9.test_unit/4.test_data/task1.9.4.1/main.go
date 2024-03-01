package main

func main() {

}

func average(xs []float64) float64 {
	var sum float64
	for i:=0; i<len(xs); i++ {
		sum = sum + xs[i]
	}
	return sum/float64(len(xs))
}
