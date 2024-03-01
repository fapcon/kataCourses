package main

import "fmt"

func Sum(a ...int) int {
	sum := 0
	for _, a := range a {
		sum += a
	}
	return sum
}

func Mul(a ...int) int {
	mul := 1
	for _, a := range a {
		mul *= a
	}
	return mul
}

func Sub(a ...int) int {
	var s []int
	sub := a[0]
	for _, a := range a {
		s = append(s, a)
	}
	for i:= 1; i<len(s); i++ {
		sub = sub-s[i]
	}
	return sub
}

func MathOperate(op func(a ...int) int, a ...int) int {
	x := op(a...)
	return x
}

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))  // Output: 5
	fmt.Println(MathOperate(Mul, 1, 7, 3))  // Output: 21
	fmt.Println(MathOperate(Sub, 13, 2, 3)) // Output: 8
}
