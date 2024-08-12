package main

import (
	"encoding/json"
	"fmt"
)

var jsonPersons string = `[
	{"firstName": "Andrew", "age": 34},
	{"firstName": "Anna", "age": 23}
]`

type Person struct {
	FirstName string
	Age       uint8
}

func main() {
	var people []Person
	err := json.Unmarshal([]byte(jsonPersons), &people)
	if err != nil {
		fmt.Println(err)
	}
	for _, person := range people {
		fmt.Printf("Name: %s. Age: %d\n", person.FirstName, person.Age)
	}
}
