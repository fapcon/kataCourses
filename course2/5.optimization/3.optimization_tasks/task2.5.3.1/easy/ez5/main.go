package main

import "fmt"

func main() {
	fmt.Println(numberOfMatches(2))
}

func numberOfMatches(n int) int {
	if n <= 1 || n >= 200 {
		return 0
	}
	var m, sum int
	sum = n / 2
	for n >= 2 {
		if n%2 == 0 {
			n = n / 2
			m = n / 2
		} else if n%2 != 0 {
			n = n/2 + 1
			m = n / 2
		}
		sum = sum + m
	}
	return sum
}
