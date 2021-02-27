package main

import "fmt"

type setS struct {
	a bool
	b bool
}

type setM map[int]bool

func main() {
	setS1 := setS{
		a: true,
		b: false,
	}
	fmt.Println("setS1", setS1)

	setS2 := setM{
		10: true,
		20: false,
	}
	fmt.Println("setS2", setS2)
}
