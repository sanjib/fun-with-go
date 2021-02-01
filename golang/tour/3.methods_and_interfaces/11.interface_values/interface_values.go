package main

import "fmt"

// I is an interface
type I interface {
	M()
}

// T is a struct
type T struct {
	S string
}

// M is implemented by T
func (t T) M() {
	fmt.Println(t.S)
}

// F is of type float64
type F float64

// M is implemented by F
func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var t I = T{"hello"}
	var f I = F(3.14)

	describe(t)
	describe(f)
}
