package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

func (c Cat) Speak() {
	fmt.Println("meow")
}

func (d Dog) Speak() {
	fmt.Println("woof")
}

func main() {
	d1 := Dog{}
	d1.Speak()

	c1 := Cat{}
	c1.Speak()
}
