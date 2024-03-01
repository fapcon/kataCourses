
package main

import (
"fmt"
"strconv"
"strings"
)

func getIntMaxValue(in8 int8, in16 int16, in32 int32, in64 int64) (int8, int16, int32, int64) {
	fmt.Println(getBits(in8)-1)
	a := in8 << (getBits(in8)-1)-1
	b := in16 << (getBits(in16)-1)-1
	c := in32 << (getBits(in32)-1)-1
	d := in64 << (getBits(in64)-1)-1
	return  a, b, c, d
}

// Если значение не нулевое, вернуть максимальное значение для типа
func getUintMaxValue(uin8 uint8, uin16 uint16, uin32 uint32, uin64 uint64) (uint8, uint16, uint32, uint64) {
	var a uint8
	var b uint16
	var c uint32
	var d uint64
	if uin8 != 0 {
		a = (uin8 << getBits(uin8))-1
	}
	if uin16 != 0 {
		b = (uin16 << getBits(uin16))-1
	}
	if uin32 != 0 {
		c = (uin32 << getBits(uin32))-1
	}
	if uin64 != 0 {
		d = (uin64 << getBits(uin64))-1
	}
	return a, b, c, d
}

func getBits(v interface{}) int {
	rawType := fmt.Sprintf("%T", v)
	typeBits := strings.Split(rawType, "t")[1]
	bits, _ := strconv.Atoi(typeBits)

	return bits
}

// Задача 2
// Функция при передаче одинаковых переменных возвращает true используя bitwise операции
func isEqual(a,b int) bool {
	if a^b==0 {
		return true
	}
	return false
}
// Задача 3
// Написать функцию которая принимает указатель на переменную и увеличивает ее значение на 1, с помощью инкремента
func inc(a *int) {
	*a++
}

func main() {
	fmt.Println(getIntMaxValue(1, 1, 1, 1))
	fmt.Println(getUintMaxValue(1, 1, 1, 1))
}
