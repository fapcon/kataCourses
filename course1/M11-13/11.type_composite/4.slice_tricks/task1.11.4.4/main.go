package main

func main() {

}

func RemoveIDX(xs []int, idx int) []int {
	if idx > len(xs) {
		return xs
	}
	xs = append(xs[:idx], xs[idx+1:]...)
	return xs
}
