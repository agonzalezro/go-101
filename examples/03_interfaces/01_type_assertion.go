package main

import "fmt"

func output(whatever interface{}) {
	switch v := whatever.(type) {
	case string:
		fmt.Println("A String: ", v)
	case int:
		fmt.Println("Some int: ", v)
	default:
		fmt.Println("I don't know what type is this one: ", v)
	}
}

func main() {
	output("this is some text")
	output(1)
	output(1.1)
}
