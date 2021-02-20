package main

import "fmt"

func main() {
	var td1 *struct {
		a *[7]int
	}
	for i, _ := range td1.a {
		fmt.Printf("i %d\n", i)
	}

	var td2 struct {
		a [7]int
	}
	for i, s := range td2.a {
		fmt.Printf("i %d, s %#v\n", i, s)
	}

	var a [10]string
	for i, s := range a {
		fmt.Printf("i %d, s %#v\n", i, s)
	}
}
