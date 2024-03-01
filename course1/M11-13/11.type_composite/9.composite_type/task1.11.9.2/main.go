package main

import "fmt"

type TVer interface {
	SamsungTV
	LgTV
	switchOFF()
	GetStatus()
	GetModel()
}

func (a *SamsungTV) switchOFF() bool {
	a.status = false
	return true
}

func (a *LgTV) switchOFF() bool {
	a.status = false
	return true
}

func (a *SamsungTV) switchOn() bool {
	a.status = true
	return true
}

func (a *LgTV) switchOn() bool {
	a.status = true
	return true
}

func (a *SamsungTV) GetStatus() bool {
	return a.status
}

func (a *LgTV) GetStatus() bool {
	return a.status
}

func (a *LgTV) GetModel() string {
	return a.model
}

func (a *SamsungTV) GetModel() string {
	return a.model
}

func (a *SamsungTV) SamsungHub() string {
	return "SamsungHub"
}

func (a *LgTV) LGHub() string {
	return "LGHub"
}

type SamsungTV struct {
	status bool
	model  string
}

type LgTV struct {
	status bool
	model  string
}

func main() {
	tv := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	fmt.Println(tv.GetStatus())  // true
	fmt.Println(tv.GetModel())   // Samsung XL-100500
	fmt.Println(tv.SamsungHub()) // SamsungHub
	fmt.Println(tv.switchOFF())  // true
	fmt.Println(tv.GetStatus())  // false
	fmt.Println(tv.switchOn())   // true
	fmt.Println(tv.GetStatus())  // true
	fmt.Println(tv.switchOFF())  // true
	fmt.Println(tv.GetStatus())  // false
}
