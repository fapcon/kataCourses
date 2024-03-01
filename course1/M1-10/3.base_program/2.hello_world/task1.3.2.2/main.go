package main

import "fmt"

var s, a, b string

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(SecondString())
	fmt.Println(ThirdString())
}
func HelloWorld() string {
	s = "Hello world!"
	return s
}
func SecondString() string {
	a = "This is second line!"
	return a
}
func ThirdString() string {
	b = "This is third line!"
	return b
}