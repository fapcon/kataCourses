package main

func main() {

}
func appendInt(xs []int, x ...int) []int {
	res := append(xs, x...)
	return res
}
