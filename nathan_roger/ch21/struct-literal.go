package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	a := location{lat: 1, long: 354}
	fmt.Printf("a %+v\n", a)

	b := location{lat: 4, long: 135}
	fmt.Printf("b %#v\n", b)

	c := a
	c.lat = 14
	fmt.Printf("c %+v\n", c)
	fmt.Printf("a %+v\n", a)
}
