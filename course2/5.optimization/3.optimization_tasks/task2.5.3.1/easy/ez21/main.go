package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	var sum int
	mul := 1
	s := strconv.Itoa(n)
	sl := strings.Split(s, "")
	res := make([]int, len(sl))
	for i := 0; i < len(sl); i++ {
		res[i], _ = strconv.Atoi(sl[i])
		sum = sum + res[i]
		mul = mul * res[i]
	}
	return mul - sum
}
