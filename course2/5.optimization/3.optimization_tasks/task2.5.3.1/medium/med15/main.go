package main

import "fmt"

func main() {
	fmt.Println(minPartitions("82734"))
}

func minPartitions(n string) int {
	maxDigit := 0
	for i := 0; i < len(n); i++ {
		digit := int(n[i] - '0')
		if digit > maxDigit {
			maxDigit = digit
		}
	}
	return maxDigit
}
