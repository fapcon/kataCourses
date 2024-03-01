package main

import "fmt"

func main() {
	p := Constructor(1, 1, 0)
	fmt.Println(p.AddCar(1))
	fmt.Println(p.AddCar(2))
	fmt.Println(p.AddCar(3))
	fmt.Println(p.AddCar(1))
}

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 {
		if p.big > 0 {
			p.big--
			return true
		} else {
			return false
		}
	}
	if carType == 2 {
		if p.medium > 0 {
			p.medium--
			return true
		} else {
			return false
		}
	}
	if carType == 3 {
		if p.small > 0 {
			p.small--
			return true
		} else {
			return false
		}
	}
	return false
}
