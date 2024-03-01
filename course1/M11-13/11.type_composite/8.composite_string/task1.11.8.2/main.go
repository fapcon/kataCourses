package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func getStringHeader(s string) reflect.StringHeader {
	gostr := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	return gostr
}

func main() {
	s := "Hello, World!"
	header := getStringHeader(s)
	fmt.Printf("Data: %v\n", header.Data)
	fmt.Printf("Len: %v\n", header.Len)
}
