package main

import "fmt"

func main() {
	xs := [8]int{1, 2, 3, 4, 5, 9, 8, 7}
	xsf := [8]float64{1.1, 2.2, 3.3, 4.4, 5.5, 9.9, 8.8, 7.7}
	fmt.Println("Sorted Int Array (Descending):", sortDescInt(xs))
	fmt.Println("Sorted Int Array (Ascending):", sortAscInt(xs))
	fmt.Println("Sorted Float Array (Descending):", sortDescFloat(xsf))
	fmt.Println("Sorted Float Array (Ascending):", sortAscFloat(xsf))
}

func sortDescInt(a [8]int) [8]int {
	for i := 0; i < 7; i++ {
		for j := i + 1; j < 8; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func sortAscInt(a [8]int) [8]int {
	for i := 0; i < 7; i++ {
		for j := i + 1; j < 8; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func sortDescFloat(a [8]float64) [8]float64 {
	for i := 0; i < 7; i++ {
		for j := i + 1; j < 8; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func sortAscFloat(a [8]float64) [8]float64 {
	for i := 0; i < 7; i++ {
		for j := i + 1; j < 8; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
