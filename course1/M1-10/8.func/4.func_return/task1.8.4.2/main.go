package main

func main() {

}

func FindMaxAndMin(n ...int) (int, int) {
	var s []int
	var min, max int
	if len(n) == 0 {
		return 0, 0
	}
	for _, n := range n {
		s = append(s, n)
	}
	min = s[0]
	max = s[0]
	for i:=0; i<len(s); i++ {
		if s[i] < min {
			min=s[i]
		}
		if s[i] > max {
			max = s[i]
		}
	}
	return max, min
}
