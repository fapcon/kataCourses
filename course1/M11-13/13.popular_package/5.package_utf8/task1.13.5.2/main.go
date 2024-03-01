package main

import (
	"fmt"
	"strings"
)

func main() {
	result := countRussianLetters("Привет, мир!")
	for key, value := range result {
		fmt.Printf("%c: %d ", key, value) // в: 1 е: 1 т: 1 м: 1 п: 1 р: 2 и: 2
	}
}

var d = []rune{1072, 1073, 1074, 1075, 1076, 1077, 1078, 1079, 1080, 1081, 1082, 1083, 1084, 1085, 1086, 1087, 1088, 1089, 1090, 1091, 1092, 1093, 1094, 1095, 1096, 1097, 1098, 1099, 1100, 1101, 1102, 1103}

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	r := strings.ToLower(s)
	for _, char := range r {
		for _, val := range d {
			if char == val {
				counts[char]++
			}
		}
	}
	return counts
}

//1072, 1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1086,1087,1088,1089,1090,1091,1092,1093,1094,1095,1096,1097,1098,1099,1100,1101,1102,1103
