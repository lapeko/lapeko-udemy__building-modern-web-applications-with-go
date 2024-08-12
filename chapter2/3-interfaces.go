package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Cat struct {
	Name string
}

func (cat Cat) Says() string {
	return "Meow"
}

func (cat Cat) NumberOfLegs() int {
	return 4
}

type Dog struct {
	Name  string
	Color string
}

func (cat *Dog) Says() string {
	return "Woof"
}

func (cat *Dog) NumberOfLegs() int {
	return 4
}

func PrintInfo(animal Animal) {
	fmt.Printf("Animal says: %s. It has %d lags.\n", animal.Says(), animal.NumberOfLegs())
}

func main() {
	cat := Cat{
		Name: "Tom",
	}
	PrintInfo(&cat)
}
