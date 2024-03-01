package main

import (
	"strconv"
	"strings"
)

func main() {

}

func UserInfo(name string, age int, cities ...string) string {
	var c []string
	var a,res string

	for _, cities := range cities {
		c = append(c, cities)
		a = a+cities+", "
	}
	res  = strings.TrimSuffix(a, ", ")



	return "Имя: " + name + ", возраст: " + strconv.Itoa(age) + ", города: "+res

}