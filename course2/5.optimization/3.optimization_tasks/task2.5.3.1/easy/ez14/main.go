package main

import "fmt"

func main() {
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
}

func maximumWealth(accounts [][]int) int {
	accs := make([]int, len(accounts))
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			accs[i] = accs[i] + accounts[i][j]
		}
	}
	for i := 0; i < len(accs)-1; i++ {
		for j := 0; j < len(accs)-i-1; j++ {
			if accs[j] > accs[j+1] {
				accs[j], accs[j+1] = accs[j+1], accs[j]
			}
		}
	}
	return accs[len(accs)-1]
}
