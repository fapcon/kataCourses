package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbbAAa"))
}

func numJewelsInStones(jewels string, stones string) int {
	s1 := strings.Split(jewels, "")
	s2 := strings.Split(stones, "")
	var res int
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(jewels); j++ {
			if s2[i] == s1[j] {
				res++
			}
		}
	}
	return res
}
