package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("111.222.333.444"))
}

func defangIPaddr(address string) string {
	s := strings.Split(address, ".")
	res := strings.Join(s, "[.]")
	return res
}
