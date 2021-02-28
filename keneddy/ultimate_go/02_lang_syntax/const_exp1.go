package main

import (
	"fmt"
)

func main() {
	//// typed
	//const ui = 12345
	//const uf = 3.141592
	//
	//// untyped
	//const ti int = 123456
	//const tf float64 = 3.141592

	var answer = 3 * 0.333
	fmt.Printf("%T %#[1]v\n", answer)

	const third = 1 / 3.0
	fmt.Printf("%T %#[1]v\n", third)

	const zero = 1 / 3
	fmt.Printf("%T %#[1]v\n", zero)

	const one int8 = 1
	const two = 2 * one
	fmt.Printf("%T %#[1]v\n", two)

	const (
		maxInt = 9223372036854775807
		bigger = 9223372036854775807123456789
	)
	fmt.Printf("%T %#[1]v\n", maxInt)

}
