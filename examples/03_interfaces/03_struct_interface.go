package main

import "fmt"

type Dev interface {
	knowsGo() bool
}

type (
	JavaDev struct{ goSkill bool }
)

func (dev JavaDev) knowsGo() bool {
	return dev.goSkill
}

// Show what would happen without the *
func (dev *JavaDev) learnGo() {
	dev.goSkill = true
}

func CanBeAPaaSero(dev Dev) bool {
	return dev.knowsGo()
}

func main() {
	dev := JavaDev{}
	dev.learnGo()
	fmt.Println(CanBeAPaaSero(dev))
}
