package main

import (
	"math/big"
	"strings"
)

func main() {

}
func Add(a, b int) *int {
	var p *int
	r := a+b
   p  = &r
   return p
}
func Max(numbers []int) *int {
	var p *int
	max := numbers[0]
for i:=0; i < len(numbers); i++ {

	if numbers[i] > max {
		max = numbers[i]
	}
}
p  = &max
return p
}
func IsPrime(number int) *bool {
	var s bool
	var p *bool
	if big.NewInt(int64(number)).ProbablyPrime(0) {
		s = true
	} else {
		s = false
	}
	p = &s
return p
}
func ConcatenateStrings(strs []string) *string {
	var res string
	var p *string
res = strings.Join(strs, "")
p  = &res
return p

}

