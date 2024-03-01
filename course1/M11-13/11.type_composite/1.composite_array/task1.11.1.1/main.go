package main

func main() {

}


func sum(xs [8]int) int {
	sum := 0
for _, xs := range xs {
	sum = sum+xs
}
return sum
}

func average(xs [8]int) float64 {
	sum := 0
	for _, xs := range xs {
		sum = sum+xs
	}
	return float64(sum)/8
}

func averageFloat(xs [8]float64) float64 {
	sum := 0.0
	for _, xs := range xs {
		sum = sum+xs
	}
	return sum/8
}

func reverse(xs [8]int) [8]int {
	var x [8]int
	temp := 0
for i:=0; i<8; i++ {
	temp = xs[i]
    x[i]=xs[7-i]
    x[7-i]=temp
}
return x
}
