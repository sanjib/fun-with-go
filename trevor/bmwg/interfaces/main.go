package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog1 := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}
	PrintInfo(&dog1)

	gorilla1 := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla1)
}

func (d *Dog) Says() string {
	return "woof"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "raaaah"
}
func (d *Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	fmt.Printf("the animal %s says %s and has %d legs\n", reflect.TypeOf(a), a.Says(), a.NumberOfLegs())
}
