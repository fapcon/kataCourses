package main

import (
	"reflect"
	"testing"
)

func TestGeneralSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "c1",
			args: args{arr: []int{5, 2, 3, 4, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "c2",
			args: args{arr: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeneralSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneralSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "c1",
			args: args{arr: []int{5, 2, 3, 4, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertionSort(tt.args.arr)
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "c1",
			args: args{left: []int{5, 2, 3}, right: []int{4, 1}},
			want: []int{4, 1, 5, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "c1",
			args: args{arr: []int{5, 2, 3, 4, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "c1",
			args: args{arr: []int{5, 2, 3, 4, 1}, low: 4, high: 0},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.arr, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "c1",
			args: args{arr: []int{5, 2, 3, 4, 1}, low: 4, high: 0},
			want: []int{5, 2, 3, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.arr, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "c1",
			args: args{arr: []int{5, 2, 3, 4, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectionSort(tt.args.arr)
		})
	}
}
