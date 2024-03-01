package main

import "fmt"

var s string
func main() {
	fmt.Println(HelloWorld())
}
func HelloWorld() string {
	s = "Hello world!"
	return s
}