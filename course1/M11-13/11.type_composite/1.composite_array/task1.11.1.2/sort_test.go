package main

import (
	"fmt"
	"testing"
)

func TestSortDesc(t *testing.T) {
	result := sortDescInt([8]int{1,2,3,4,5,7,8,9})
	a := [8]int{1,2,3,4,5,7,8,9}
	for i:=0; i<7; i++ {
		for j:=i+1;j<8;j++{
			if a[i]<a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	if result != a {
		fmt.Errorf("expected: %v, got: %v", a, result)
	}
}

func TestSortAsc(t *testing.T) {
	result := sortAscInt([8]int{9,8,7,5,4,3,2,1})
	a := [8]int{9,8,7,5,4,3,2,1}
	for i:=0; i<7; i++ {
		for j:=i+1;j<8;j++{
			if a[i]>a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	if result != a {
		fmt.Errorf("expected: %v, got: %v", a, result)
	}
}

func TestSortDescc(t *testing.T) {
	result := sortDescFloat([8]float64{1.1,2.2,3.3,4.4,5.5,7.7,8.8,9.9})
	a := [8]float64{1.1,2.2,3.3,4.4,5.5,7.7,8.8,9.9}
	for i:=0; i<7; i++ {
		for j:=i+1;j<8;j++{
			if a[i]<a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	if result != a {
		fmt.Errorf("expected: %v, got: %v", a, result)
	}
}

func TestSortAscc(t *testing.T) {
	result := sortAscFloat([8]float64{9.9,8.8,7.7,5.5,4.4,3.3,2.2,1.1})
	a := [8]float64{9.9,8.8,7.7,5.5,4.4,3.3,2.2,1.1}
	for i:=0; i<7; i++ {
		for j:=i+1;j<8;j++{
			if a[i]>a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	if result != a {
		fmt.Errorf("expected: %v, got: %v", a, result)
	}
}
