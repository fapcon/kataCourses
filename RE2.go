// Написать программу fizz buzz

// Если число делится на 3, то вывести fizz
// Если число делится на 5, то вывести buzz
// Если число делится на 3 и на 5, то вывести fizz buzz

// Покрыть табличными тестами



// Задача 2
// Написать бенчмарк для Fibonacci, двух реализаций, при помощи рекурсии и формулы Бине


// Задача 3
// Написать функцию округления float, используя math.Round и math.Pow



package main

import (
	"math"
)

func main() {

}
func fizzbuzz(a int) string {
	if a%3==0 {
		return "fizz"
	}
	if a%5==0 {
		return "buzz"
	}
	if (a%3==0) && (a%5==0) {
		return "fizz buzz"
	}
	return ""
}


func FibonacciRe(n int) int {
	if n == 0 {
		return 0
	}
	if n ==1 {
		return 1
	} else {
		return FibonacciRe(n-1)+FibonacciRe(n-2)
	}
}

func FibonacciBinet(n int) int {
	phi := (1 + math.Sqrt(5)) / 2
	return int(math.Round(math.Pow(phi, float64(n)) / math.Sqrt(5)))
}



func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}


