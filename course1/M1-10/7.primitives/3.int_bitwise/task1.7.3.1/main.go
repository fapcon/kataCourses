package main

import "fmt"

func main() {
fmt.Println("a & b =", bitwiseAnd(3, 5))
	fmt.Println("a | b =", bitwiseOr(3, 5))
	fmt.Println("a ^ b =", bitwiseXor(3, 5))
	fmt.Println("a << b =", bitwiseLeftShift(3, 5))
	fmt.Println("a >> b =", bitwiseRightShift(3, 5))
}

func bitwiseAnd(a, b int) int {
	res := a & b
	return res
}
func bitwiseOr(a, b int) int {
	res := a | b
	return res
}
func bitwiseXor(a, b int) int {
	res := a ^ b
	return res
}
func bitwiseLeftShift(a, b int) int {
	res := a << b
	return res
}
func bitwiseRightShift(a, b int) int {
	res := a >> b
	return res
}
