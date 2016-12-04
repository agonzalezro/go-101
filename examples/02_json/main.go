package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

func main() {
	var sentPeople = []Person{Person{"Juan"}, Person{"Manolo"}}
	payload, err := json.Marshal(sentPeople)
	if err != nil {
		panic(err) // Don't do this at home
	}

	fmt.Println(string(payload))

	var receivedPeople []Person
	// emulatedBody := strings.NewReader(string(payload))
	// if err := json.NewDecoder(emulatedBody).Decode(&receivedPeople); err != nil {
	if err := json.Unmarshal(payload, &receivedPeople); err != nil {
		panic(err)
	}

	for _, p := range receivedPeople {
		fmt.Println(p.Name)
	}
}
