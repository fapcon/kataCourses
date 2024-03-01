package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return (n) * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// выдает true, если реализация быстрее и false, если медленнее
func compareWhichFactorialIsFaster() map[string]bool {
	num := 2

	startR := time.Now()
	factorialRecursive(num)
	endR := time.Now()

	startI := time.Now()
	factorialIterative(num)
	endI := time.Now()

	timeR := endR.Sub(startR)
	timeI := endI.Sub(startI)
	fmt.Printf("timeR = %v\n", timeR)
	fmt.Printf("timeI = %v\n", timeI)
	result := make(map[string]bool)
	if timeR < timeI {
		result["recursive"] = true
		result["iterative"] = false
	} else {
		result["recursive"] = false
		result["iterative"] = true
	}

	return result
}

func main() {
	fmt.Println(factorialIterative(15))
	fmt.Println(factorialRecursive(15))
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Go OS/Arch:", runtime.GOOS, "/", runtime.GOARCH)

	fmt.Println("Which factorial is faster?")
	fmt.Println(compareWhichFactorialIsFaster())
}
