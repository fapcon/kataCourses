package main

import (
	"strings"
)

func main() {

}

func ConcatenateStrings(sep string, stringss ...string) string {
	var c []string
	var a,b string
	for _, stringss := range stringss {
		c = append(c, stringss)
	}

	for i:=0; i<len(c); i++ {

		if i%2 == 0{
			a = a+c[i]+sep

		} else {
			b = b+c[i]+sep

		}
	}
	res1 := strings.TrimSuffix(a, sep)
	res2 := strings.TrimSuffix(b, sep)
	ress := "even: "+res1+", odd: "+res2
	return ress
}
