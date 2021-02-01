package main

import "fmt"

// I interface
type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i)
	i.M()
}
