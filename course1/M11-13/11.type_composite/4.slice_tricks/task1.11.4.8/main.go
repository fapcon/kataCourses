package main

import "fmt"

func main() {
	fmt.Println(Shift([]int{111, 2, 3, 4, 5, 66}))
}

func Shift(xs []int) (int, []int) {
	if len(xs) == 0 {
		return 0, []int{}
	}
	res1 := xs[0]
	temp := xs[len(xs)-1]
	res := make([]int, len(xs))
	for i, val := range xs {
		res[i] = val
	}
	for i := len(res) - 1; i > 0; i-- {
		res[i] = res[i-1]
	}
	res[0] = temp
	return res1, res

}
