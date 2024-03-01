package main

import "fmt"

func main() {
	fmt.Println(interpret("G()()()()(al)"))
}

func interpret(command string) string {
	var res string
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			res += "G"
		} else if command[i] == '(' && command[i+1] == ')' {
			res += "o"
		} else if command[i] == '(' && command[i+1] == 'a' {
			res += "al"
		}
	}
	return res
}
