package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReplaceSymbols("123412341234", '1', '4'))
}

func ReplaceSymbols(str string, old, new rune) string {
	/*
		r := []rune(str)
		for _, val := range r {
			if val == old {
				val = new
			}
		}
		return string(r)
	*/
	var res string
	sl := strings.Split(str, "")
	for i := range sl {
		if strings.ContainsRune(sl[i], old) == true {
			sl[i] = string(new)
		}
	}
	for _, val := range sl {
		res = res + val
	}
	return res
}
