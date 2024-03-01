package main

import (
	"fmt"
	"testing"
)

func TestGetSubSLice(t *testing.T) {
	var xs =[]int{1,2,3,4,5,6,7,8,9}
	result := getSubSlice(xs, 2, 6)
    exp := xs[2:6]
	for i:=0;i<4;i++ {
		if result[i] != exp[i] {
			fmt.Errorf("expected: %v, got: %v", exp, result)
	}
	}
}
