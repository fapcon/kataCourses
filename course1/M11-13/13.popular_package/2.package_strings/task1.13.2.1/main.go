package main

import (
	"fmt"
	"strings"
)

func CountWordsInText(txt string, words []string) map[string]int {
	// Ваш код для подсчета слов в тексте
	res := make(map[string]int)
	r := strings.ToLower(txt)
	s := strings.Fields(r)
	for _, val := range s {
		for _, lav := range words {
			if val == lav {
				res[val]++
			}
		}
	}
	return res
}

func main() {
	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. 
        Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. 
        Praesent et diam eget libero egestas mattis sit amet vitae augue.`
	words := []string{"sit", "amet", "lorem"}

	result := CountWordsInText(txt, words)

	fmt.Println(result) // map[amet:2 lorem:1 sit:3]
}
