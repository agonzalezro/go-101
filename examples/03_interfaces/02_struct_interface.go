package main

import "fmt"

type Dev interface {
	knowsGo() bool
}

type (
	GoDev   struct{}
	JavaDev struct{}
)

func (GoDev) knowsGo() bool {
	return true
}

func (JavaDev) knowsGo() bool {
	return false
}

func CanBeAPaaSero(dev Dev) bool {
	return dev.knowsGo()
}

func main() {
	fmt.Println(CanBeAPaaSero(GoDev{}), CanBeAPaaSero(JavaDev{}))
}
