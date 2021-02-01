package main

import (
	"fmt"
	"math"
)

// Abser is an interface for absolute functions
type Abser interface {
	Abs() float64
}

// Vertex holds 2D cartesian coords
type Vertex struct {
	X, Y float64
}

// Abs returns the absolute value of Vertex, implements Abser
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MyFloat is of type float64
type MyFloat float64

// Abs returns the absolute value of MyFloat, implements Abser
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())
}
