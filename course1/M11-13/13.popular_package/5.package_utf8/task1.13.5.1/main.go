package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countUniqueUTF8Chars("Hello, 世界!"))
}
func countUniqueUTF8Chars(s string) int {
	r := strings.Split(s, "")
	m := make(map[string]int)
	for _, val := range r {
		m[val]++
	}
	return len(m)
}
