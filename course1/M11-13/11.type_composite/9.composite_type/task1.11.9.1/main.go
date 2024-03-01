package main

import "fmt"

type OrganicWorld interface {
	Cell
	NonCellular
	Live()
}

type Cell interface {
	Prokaryote
	Eukaryote
	Grow()
	Divide() Cell
}

type Prokaryote interface {
	Bacteria
	Archaea
	ProduceToxins()
}

type Bacteria interface {
}
type Archaea interface {
}

type Eukaryote interface {
	Animal
	Fungus
	Plant
	CloneGenome()
}
type Animal interface {
	Move()
	Eat()
}

type Plant interface {
}

type Fungus interface {
}

type NonCellular interface {
	Virus
	Replicate() NonCellular
}

type Virus interface {
	Infect()
}

type AnimalCell struct {
	AnimalEukaryote
}

func (c *AnimalCell) Grow() {

}

func (c *AnimalCell) Divide() Cell {
	return c
}

func (c *AnimalCell) Live() {

}

type AnimalEukaryote struct {
	Cell
}

func (e *AnimalEukaryote) CloneGenome() {
}

type AnimalCat struct {
	AnimalEukaryote
}

func (a AnimalCat) ProduceToxins() {

}

func (a AnimalCat) Move() {

}

func (a AnimalCat) Eat() {

}

func (a AnimalCat) CloneGenome() {

}

func (a AnimalCat) Grow() {

}

func (a AnimalCat) Divide() Cell {
	return a
}

func (a AnimalCat) Infect() {

}

func (a AnimalCat) Live() {

}

func (a AnimalCat) Replicate() NonCellular {
	return a
}

type InfluenzaVirus struct {
	Virus
}

func (i InfluenzaVirus) Infect() {
}

func (i InfluenzaVirus) Replicate() NonCellular {
	return i
}

func main() {

	var cell Cell = &AnimalCat{}
	cell.Grow()
	newCell := cell.Divide()
	fmt.Println(newCell)

	var nonCell NonCellular = &InfluenzaVirus{}
	nonCell.Infect()
	newNonCell := nonCell.Replicate()
	fmt.Println(newNonCell)
}
