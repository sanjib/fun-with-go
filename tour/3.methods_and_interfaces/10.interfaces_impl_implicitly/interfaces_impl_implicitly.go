package main

import "fmt"

// I is an interface for M()
type I interface {
	M()
}

// T holds S string
type T struct {
	S string
}

// M is an implementation of I by T
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
