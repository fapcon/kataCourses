package main

func main() {

}

func Factorial(n int) int {
	var res int
	res = 1
	if n == 0 {
		return res
	} else {
		for i := 1; i <= n; i++ {
			res = res * i
		}
		return res
	}
}

