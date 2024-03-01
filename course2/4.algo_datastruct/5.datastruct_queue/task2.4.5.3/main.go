package main

import (
	"container/list"
	"fmt"
)

type Car struct {
	LicensePlate string
}

type ParkingLot struct {
	space *list.List
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{space: list.New()}
}

func (p *ParkingLot) Park(c Car) {
	p.space.PushBack(c)
	fmt.Printf("Автомобиль [%v] припаркован\n", c.LicensePlate)
}

func (p *ParkingLot) Leave() {
	if p.space.Len() == 0 {
		fmt.Println("Парковка пуста.")
	} else {
		for p.space.Len() > 0 {
			element := p.space.Front()
			p.space.Remove(element)
			fmt.Printf("Автомобиль [%v] покинул парковку\n", element.Value)
		}
	}
}

func main() {
	parkingLot := NewParkingLot()
	parkingLot.Park(Car{LicensePlate: "ABC-123"})
	parkingLot.Park(Car{LicensePlate: "XYZ-789"})
	fmt.Println(parkingLot.space)
	parkingLot.Leave()
	parkingLot.Leave()
	parkingLot.Leave()
	fmt.Println(parkingLot.space)
	parkingLot.Park(Car{LicensePlate: "XYZ-789"})
	fmt.Println(parkingLot.space)

}
