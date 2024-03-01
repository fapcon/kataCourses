package main

import "fmt"

func GeneralSort(arr []int) []int {
	if len(arr) < 10 {
		res := mergeSort(arr)
		return res
	} else if len(arr) < 100 {
		res := quickSort(arr, 0, len(arr)-1)
		return res
	} else if len(arr) < 1000 {
		insertionSort(arr)
		return arr
	} else {
		selectionSort(arr)
		return arr
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original: ", data)

	sortedData := mergeSort(data)
	fmt.Println("Sorted by Merge Sort: ", sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(data)
	fmt.Println("Sorted by Insertion Sort: ", data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(data)
	fmt.Println("Sorted by Selection Sort: ", data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = quickSort(data, 0, len(data)-1)
	fmt.Println("Sorted by Quicksort: ", sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = GeneralSort(data)
	fmt.Println("Sorted by GeneralSort: ", sortedData)
}
