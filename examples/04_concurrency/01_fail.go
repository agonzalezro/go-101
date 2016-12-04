package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
	// Something is going to go wrong...
}
