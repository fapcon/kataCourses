package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func binaryStringToFloat(binary string) float32 {
var number uint32
var i int64
i = int64(number)
// Преобразование строки в двоичной системе в целочисленное представление
	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
// Преобразование целочисленного представления в число с плавающей точкой
floatNumber := *(*float32)(unsafe.Pointer(&i))
return floatNumber
}

func main() {

}
