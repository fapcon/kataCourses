package main

import (
	"fmt"
	"unsafe"
)

func uint8ToInt8(num uint8) int8 {
	res := *(*int8)(unsafe.Pointer(&num))
	return res
}

func main() {
	fmt.Println(uint8ToInt8(128)) // Вывод: -128
}
