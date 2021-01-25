package main

import "fmt"

func main() {
	a := 'A'
	fmt.Println(a)

	b := "\x41"
	fmt.Println(b)

	c := 65
	fmt.Println(c)

	x := "ABC"
	fmt.Printf("%#v %v\n", x[0], x[0])
	fmt.Printf("%#v %v\n", x[1], x[1])
	fmt.Printf("%#v %v\n", x[2], x[2])
}
