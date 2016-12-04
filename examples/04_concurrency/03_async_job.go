package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func doit(i int, done chan bool) {
	// Change this number to reach the timeout
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Println(i)
	done <- true
}

func main() {
	rand.Seed(42)
	done := make(chan bool, 1)

	for i := 0; i < 10; i++ {
		go doit(i, done)
	}

	for i := 0; i < 10; i++ {
		select {
		case <-done:
			continue
		case <-time.After(time.Second * 3):
			log.Println("Too... much... waiting...")
			return
		}
	}
}
