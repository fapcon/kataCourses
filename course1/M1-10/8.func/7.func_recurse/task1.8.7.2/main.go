package main

func main() {

}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n ==1 {
		return 1
	} else {
		return Fibonacci(n-1)+Fibonacci(n-2)
	}
}

