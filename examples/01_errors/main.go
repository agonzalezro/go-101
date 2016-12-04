package main

import (
	"errors"
	"log"
)

var NamedError = errors.New("Is this an exception?")

func breakSomething() error {
	// Return something else for the sake of the example
	// return errors.New("hi")
	return NamedError
}

func main() {
	if err := breakSomething(); err != nil {
		log.Println("Inline fail")
	}

	if err := breakSomething(); err == NamedError {
		log.Println("Named fail")
	}

	err := breakSomething()
	switch err {
	case NamedError:
		log.Println("Named fail in switch")
	default:
		log.Fatal("I didn't expect this!")
	}
}
