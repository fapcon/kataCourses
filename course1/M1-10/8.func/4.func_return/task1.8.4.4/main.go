package main

func main() {

}
func CalculateSimpleMovingAverage(period int, data ...float64) []float64 {

	var sum, mas []float64
	for _, data := range data {
		mas = append(mas, data)
	}
	var r float64
	if period >= len(mas) {
		a := 0.0
		for i:=0; i<len(mas); i++ {
			a = a+mas[i]
		}
		r = a/float64(len(mas))
		sum = append(sum, r)
		return sum
	} else {
		for j := 0; j < len(mas)-(period-1); j+=period {
			a := 0.0
			for i := 0; i+j < period+j; i++ {
				a = a + mas[i+j]
			}
			r = a / float64(period)
			sum = append(sum, r)
		}
		return sum
	}
}
