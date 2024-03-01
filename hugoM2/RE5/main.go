package main

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

type Weather interface {
	getWeather(city string) int
}

type MockWeather struct {
	mock.Mock
}

type w struct {
	value int
}

func main() {
	n := w{5}
	//m := &MockWeather{}
	fmt.Println(n.getWeather("moscow"))
	//fmt.Println(m.getWeather("moscow"))
}

func (a *w) getWeather(city string) int {
	return a.value
}

//func (_m *MockWeather) GetWeather(city string) int {
//	ret := _m.Called(city)
//	var r w
//	if f, ok := ret.Get(0)(func(string)int); ok {
//		return f(city)
//	}
//}
