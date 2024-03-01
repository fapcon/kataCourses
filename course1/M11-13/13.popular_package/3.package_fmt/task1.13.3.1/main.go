package main

import "fmt"

func generateMathString(operands []int, operator string) string {
	sum := 0
	sub := operands[0]
	var res int
	switch operator {
	case "+":
		for _, val := range operands {
			sum = sum + val
			res = sum
		}
	case "-":
		for _, val := range operands {
			sub = sub - val
			res = sub
		}
	}

	s := fmt.Sprintf("%d %s %d %s %d = %d", operands[0], operator, operands[1], operator, operands[2], res)
	return s
}

// Пример результата выполнения программы:
func main() {
	fmt.Println(generateMathString([]int{2, 4, 6}, "+")) // "2 + 4 + 6 = 12"
}
