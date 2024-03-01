package main

import "fmt"

//var Operate func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} // реализуй меня

func Operate(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}

//var Concat func(xs ...interface{}) interface{}

func Concat(xs ...interface{}) interface{} {

	var mas []string
	var res string
	for _, val := range xs {
		mas = append(mas, val.(string))
	}
	for _, val := range mas {
		res += val
	}
	return res
}

//var Sum func(xs ...interface{}) interface{} // реализуй меня для int и float64

func Sum(xs ...interface{}) interface{} {
	switch xs[0].(type) {
	case float64:
		var mas []float64
		var res float64
		for _, val := range xs {
			mas = append(mas, val.(float64))
		}
		for _, val := range mas {
			res += val
		}
		return res
	case int:
		var mas []int
		var res int
		for _, val := range xs {
			mas = append(mas, val.(int))
		}
		for _, val := range mas {
			res += val
		}
		return res
	}
	return 0
}

func main() {

	fmt.Println(Operate(Concat, "Hello, ", "World!"))  // Вывод: "Hello, World!"
	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))           // Вывод: 15
	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5)) // Вывод: 16.5
}
