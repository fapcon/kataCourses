package main

import (
	"strings"
)

func main() {

}

func mutate(a *int) {
*a = 42
}

func ReverseString(str *string) {
var s = strings.Split(*str, "")
	for i := len(s)/2-1; i >= 0; i-- {
		opp := len(s)-1-i
		s[i], s[opp] = s[opp], s[i]
	}
	*str = strings.Join(s, "")
}
