package main

import "fmt"

// I interface
type I interface {
	M()
}

// T struct
type T struct {
	S string
}

// M method
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Printf("%q\n", t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	var t *T

	i = t
	i.M()
	describe(i)

	i = &T{"hi"}
	i.M()
	describe(i)
}
