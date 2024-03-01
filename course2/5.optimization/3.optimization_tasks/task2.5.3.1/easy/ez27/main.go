package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRRLLRLL"))
}

func balancedStringSplit(s string) int {
	var a, b, c int
	for _, i := range s {
		if i == 'R' {
			a++
		} else if i == 'L' {
			b++
		}

		if a == b {
			a = 0
			b = 0
			c++
		}
	}
	return c
}
