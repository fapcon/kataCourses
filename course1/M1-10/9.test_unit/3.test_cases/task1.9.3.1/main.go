package main

func main() {
	
}

func Fibonacci(a int) int {
	var x int

	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}
	if a == 2 {
		return 1
	} else {
		s:=make([]int, a, a)
		s[0]=1
		s[1]=1
		for i:=2; i<a; i++{
			s[i] = s[i-1]+s[i-2]
			x = s[i]
		}
		return x
	}
}