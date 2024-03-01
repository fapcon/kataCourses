package main

import "fmt"

func main() {
	a := 5
	b := 10
	c := Dereference(&a)
	d := Sum(&b, &c)
	fmt.Println(c) // Output: 5
	fmt.Println(d) // Output: 15
}

func Dereference(n *int) int {
var x int = *n
return x
}

func Sum(a,b *int) int {
var sum int = *b+*a
return sum
}