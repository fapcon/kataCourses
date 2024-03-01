package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

var a, b string
var precision int32

func main() {
    fmt.Println(DecimalSum("2.333333", "3.222222"))
	fmt.Println(DecimalSubtract("2.333333", "3.222222"))
	fmt.Println(DecimalMultiply("2.333333", "3.222222"))
	fmt.Println(DecimalDivide("2.333333", "3.222222"))
	fmt.Println(DecimalRound("2.333333", 2))
	fmt.Println(DecimalGreaterThan("2.333333", "3.222222"))
	fmt.Println(DecimalLessThan("2.333333", "3.222222"))
	fmt.Println(DecimalEqual("2.333333", "3.222222"))
}
func DecimalSum(a, b string) (string, error) {
	result, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	result1, err1 := decimal.NewFromString(b)
	if err != nil {
		panic(err1)
	}
	result2 := result.Add(result1)
    return result2.String(), err
}

func DecimalSubtract(a, b string) (string, error) {
	result, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	result1, err1 := decimal.NewFromString(b)
	if err != nil {
		panic(err1)
	}
	result2 := result.Sub(result1)
	return result2.String(), err
}

func DecimalMultiply(a, b string) (string, error) {
	result, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	result1, err1 := decimal.NewFromString(b)
	if err != nil {
		panic(err1)
	}
	result2 := result.Mul(result1)
	return result2.String(), err
}

func DecimalDivide(a, b string) (string, error) {
	result, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	result1, err1 := decimal.NewFromString(b)
	if err != nil {
		panic(err1)
	}
	result2 := result.Div(result1)
	return result2.String(), err
}

func DecimalRound(a string, precision int32) (string, error) {
	result, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}

	result2 := result.Round(precision)
	return result2.String(), err
}

func DecimalGreaterThan(a, b string) (bool, error) {
	result, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	result1, err1 := decimal.NewFromString(b)
	if err != nil {
		panic(err1)
	}
	result2 := result.GreaterThan(result1)
	return result2, err
}

func DecimalLessThan(a, b string) (bool, error) {
	result, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	result1, err1 := decimal.NewFromString(b)
	if err != nil {
		panic(err1)
	}
	result2 := result.LessThan(result1)
	return result2, err
}

func DecimalEqual(a, b string) (bool, error) {
	result, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	result1, err1 := decimal.NewFromString(b)
	if err != nil {
		panic(err1)
	}
	result2 := result.Equal(result1)
	return result2, err
}
