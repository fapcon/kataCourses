package main

import (
	"fmt"
	"strconv"
)

type Mover interface {
	Move() string
	Speed() int
	MaxSpeed() int
	MinSpeed() int
}

type FastMover struct {
	BaseMover
}

func (f FastMover) Move() string {
	return "Fast mover! Moving at speed: " + strconv.Itoa(f.speed)
}

func (f FastMover) Speed() int {
	return f.speed
}

func (f FastMover) MaxSpeed() int {
	return 120
}

func (f FastMover) MinSpeed() int {
	return 10
}

type SlowMover struct {
	BaseMover
}

func (s SlowMover) Move() string {
	return "Slow mover! Moving at speed: " + strconv.Itoa(s.speed)
}

func (s SlowMover) Speed() int {
	return s.speed
}

func (s SlowMover) MaxSpeed() int {
	return 120
}

func (s SlowMover) MinSpeed() int {
	return 10
}

type BaseMover struct {
	speed int
}

func main() {
	var movers []Mover
	fm := FastMover{BaseMover{100}}
	sm := SlowMover{BaseMover{10}}
	movers = append(movers, fm, sm)

	for _, mover := range movers {
		fmt.Println(mover.Move())
		fmt.Println("Maximum speed:", mover.MaxSpeed())
		fmt.Println("Minimum speed:", mover.MinSpeed())
	}
}
