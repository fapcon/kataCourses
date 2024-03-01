package main

import "fmt"

func main() {
	fmt.Println(convertTemperature(30))
}

func convertTemperature(celsius float64) []float64 {
	if celsius < 0 || celsius > 1000 {
		return nil
	} else {
		k := celsius + 273.15
		f := celsius*1.80 + 32.00
		return []float64{k, f}
	}
}
