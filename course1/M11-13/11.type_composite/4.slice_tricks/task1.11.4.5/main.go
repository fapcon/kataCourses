package main

func main() {

}

func FilterDividers(xs []int, divider int) []int {
	var a []int
	if len(xs) == 0 {
		return []int{}
	}
	for i := 0; i < len(xs); i++ {
		if xs[i]%divider == 0 {
			a = append(a, xs[i])
		}
	}
	return a
}
