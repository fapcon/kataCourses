package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)
	myPanic(ch)
	fmt.Println(<-ch)
}

func myPanic(ch chan string) {

	defer func() {
		if r := recover(); r != nil {
			ch <- r.(string)
		}
	}()
	panic("my panic message")
}