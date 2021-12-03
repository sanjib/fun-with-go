package main

import "fmt"

type circle struct {
	radius int
	border int
}

func main() {
	x := 10
	f := 123.45

	fmt.Printf("x dec: %d\nx hex: %x\n", x, x)
	fmt.Printf("f float: %f\nf exp: %e\n", f, f)
	fmt.Printf("bool: %t\n", true)

	fmt.Printf("%[1]d %[1]d\n", 52, 40)
	fmt.Printf("%[2]d %[1]d\n", 52, 40)
	fmt.Printf("%d %#[1]x %#[1]o %#[1]b\n", 16)

	c := circle{10, 20}
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
}
