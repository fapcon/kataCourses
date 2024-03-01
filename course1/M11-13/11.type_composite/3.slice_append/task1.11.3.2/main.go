package main

func main() {

}

func appendInt(xs *[]int, x ...int) {
	*xs = append(*xs, x...)
}
