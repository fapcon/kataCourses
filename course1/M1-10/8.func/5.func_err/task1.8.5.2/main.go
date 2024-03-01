package main

import (
	"errors"
)

func main() {

}

func CheckDiscount(price, discount float64) (float64, error) {
	var mul float64
	var err error
	if discount > 50 {
		return 0, errors.New("Скидка не может превышать 50%")
	}
	mul = discount/100
	res := 100 - price*mul
	return res, err
}
