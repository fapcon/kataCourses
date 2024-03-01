package main

func main() {

}

func Cut(xs []int, start, end int) []int {
	def := []int{}
	if end > len(xs) || start > len(xs) {
		return def
	}
	if start < 0 || end < 0 {
		return def
	}
	res := xs[start : end+1]
	return res
}
