package main

import "fmt"

// Определение стека
var stack = []int{}

// Функция для добавления элемента в стек
func push(value int) {
	stack = append(stack, value)
	fmt.Println(stack)
}

// Функция для удаления и возврата последнего элемента из стека
func pop() int {
	if len(stack) == 0 {
		panic("Стек пуст")
	}
	lastIndex := len(stack) - 1
	value := stack[lastIndex]
	stack = stack[:lastIndex]
	return value
}

// Пример использования стека для операций
func main() {
	push(5)
	push(3)
	push(10)
	result := pop() + pop()
	push(result)
}
