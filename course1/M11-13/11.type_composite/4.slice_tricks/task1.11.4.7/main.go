package main

import "fmt"

func main() {

}

func Pop(xs []int) (int, []int) {
	if len(xs) == 0 {
		fmt.Printf("Значение: %d, Новый срез: %v", 0, []int{})
		return 0, []int{}
	}
	if len(xs) == 1 {
		fmt.Printf("Значение: %d, Новый срез: %v", xs[0], []int{})
		return xs[0], []int{}
	}
	fmt.Printf("Значение: %d, Новый срез: %v", xs[0], xs[1:])
	return xs[0], xs[1:]
}
