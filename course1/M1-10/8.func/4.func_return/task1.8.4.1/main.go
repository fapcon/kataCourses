package main

import (
	"fmt"
)

func main() {
	a, b:=DivideAndRemainder(6, 5)
fmt.Printf("Частное: %d, Остаток: %d", a, b)
}

func DivideAndRemainder(a, b int) (int, int) {
	if b != 0 {
		return a / b, a % b
	}
	return 0, 0
}
