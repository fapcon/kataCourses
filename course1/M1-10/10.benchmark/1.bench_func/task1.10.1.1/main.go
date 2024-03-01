package main

func main() {

}

func Fibonacci(n int) int {
	var x int

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	} else {
		s:=make([]int, n, n)
		s[0]=1
		s[1]=1
		for i:=2; i<n; i++{
			s[i] = s[i-1]+s[i-2]
			x = s[i]
		}
		return x
	}
}
