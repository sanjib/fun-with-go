package main

import "fmt"

func main() {
	v := 42
	w := 42.0
	x := "42"
	y := true
	z := .3 + .1i
	fmt.Printf("%v %T\n", v, v)
	fmt.Printf("%v %T\n", w, w)
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", z, z)
}
