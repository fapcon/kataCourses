package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countDigits(122))
}

func countDigits(num int) int {
	str := strconv.Itoa(num)
	sl := strings.Split(str, "")
	r := make([]int, len(sl))
	var count int
	for i := 0; i < len(sl); i++ {
		r[i], _ = strconv.Atoi(sl[i])
		if r[i] != 0 {
			if num%r[i] == 0 {
				count++
			}
		}
	}
	return count
}
