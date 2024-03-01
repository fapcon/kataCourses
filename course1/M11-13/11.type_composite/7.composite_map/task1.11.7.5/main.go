package main

import (
	"fmt"
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	s := strings.Split(sentence, " ")
	i := 0
	keys := make([]string, len(filter))
	for k := range filter {
		keys[i] = k
		i++
	}
	for j := 0; j < len(s)-1; j++ {
		//for k, v := range s {
		for i := 0; i < len(keys); i++ {
			if s[j] == keys[i] {

				s = append(s[:j], s[j+1:]...)

			}
		}
	}
	for i := 0; i < len(keys); i++ {
		if keys[i] == s[len(s)-1] {
			s = s[:len(s)-1]
		}
	}
	var res string
	for i := 0; i < len(s)-1; i++ {
		res = res + s[i] + " "
	}
	res = res + s[len(s)-1]
	return res
}

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	filteredSentence := filterSentence(sentence, filter)
	fmt.Println(filteredSentence)
}
