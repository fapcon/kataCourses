package main

import "fmt"

func mergeMaps(map1, map2 map[string]int) map[string]int {
	res := make(map[string]int)
	for i, v := range map1 {
		res[i] = v
	}
	for j, val := range map2 {
		res[j] = val
	}
	return res
}

func main() {
	map1 := map[string]int{"apple": 3, "banana": 2}
	map2 := map[string]int{"orange": 5, "grape": 4}
	mergedMap := mergeMaps(map1, map2)
	for key, value := range mergedMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
