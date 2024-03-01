package main

import "fmt"

type SubrectangleQueries struct {
	rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rectangle[row][col]
}

func main() {
	rectangle := [][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}}
	r := Constructor(rectangle)
	fmt.Println(r.GetValue(1, 1)) // return 1
	r.UpdateSubrectangle(0, 0, 3, 1, 10)
	fmt.Println(r)

}
