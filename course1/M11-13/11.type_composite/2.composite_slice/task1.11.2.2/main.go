package main


func main() {

}

func MaxDifference(numbers []int) int {
	if len(numbers) == 0 || len(numbers) == 1 {
		return 0
	}
	max := numbers[0]
	min := numbers[0]
	for i:=0; i<len(numbers); i++ {
		if max<numbers[i] {
			max = numbers[i]
		}
		if min>numbers[i] {
			min = numbers[i]
		}
	}
	return max-min
}
