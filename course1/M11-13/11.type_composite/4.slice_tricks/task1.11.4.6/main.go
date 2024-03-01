package main

func main() {

}
func InsertToStart(xs []int, x ...int) []int {
	var a []int
	for _, val := range x {
		a = append(a, val)
	}
	for _, s := range xs {
		a = append(a, s)
	}
	return a
}
