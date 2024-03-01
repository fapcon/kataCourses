package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(createUniqueText("qwe asd qwe asd asd qwe qwe ytr ytr"))
}

func createUniqueText(text string) string {
	var res string
	res = ""
	s := strings.Split(text, " ")
	lenres := make(map[string]int)
	counts := make(map[string]int)
	for _, v := range s {
		lenres[v]++
	}
	sl := make([]string, len(lenres))
	for _, v := range s {
		counts[v]++
		if counts[v] == 1 {
			sl = append(sl, v)
		}
	}

	for _, v := range sl {
		res = res + v + " "
	}

	k := strings.TrimSpace(res)
	return k
}
