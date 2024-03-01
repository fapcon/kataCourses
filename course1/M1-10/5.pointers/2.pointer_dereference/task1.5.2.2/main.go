package main

import (
	"math"
	"strings"
)

func main() {

}

func Factorial(n *int) int {
var a int = *n
res := 1
	for i:=a; i>=1; i-- {
		res *= i
	}
	return res
}

func isPalindrome(str *string) bool {
var s = strings.Split(*str, "")
var a int
var b = len(s)
var c = float64(b)
var res bool
for i := 0; i <= len(s)-1; i++ {
	if strings.ToLower(s[i]) == strings.ToLower(s[len(s)-1-i]) {
		a++
	}
}
if float64(a) == math.Floor(c) {
	res = true
} else {
	res = false
}
return res
}

func CountOccurrences(numbers *[]int, target *int) int {
var a int = *target
var b []int = *numbers
var res int
for i:=0; i<=len(b)-1; i++ {
	if b[i] == a {
		res++
	}
}
	return res
}

func ReverseString(str *string) string {
	var s = strings.Split(*str, "")
	for i := len(s)/2-1; i >= 0; i-- {
		opp := len(s)-1-i
		s[i], s[opp] = s[opp], s[i]
	}
	res := strings.Join(s, "")
	return res
}
