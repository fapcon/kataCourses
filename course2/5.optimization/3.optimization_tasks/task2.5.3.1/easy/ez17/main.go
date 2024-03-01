package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))
}

func mostWordsFound(sentences []string) int {
	sum := make([]int, len(sentences))
	for i := 0; i < len(sentences); i++ {
		sum[i] = len(strings.Split(sentences[i], " "))
	}
	for i := 0; i < len(sum)-1; i++ {
		for j := 0; j < len(sum)-i-1; j++ {
			if sum[j] > sum[j+1] {
				sum[j], sum[j+1] = sum[j+1], sum[j]
			}
		}
	}
	return sum[len(sum)-1]
}
