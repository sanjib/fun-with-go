package main

import (
	"fmt"
	"math"
)

// Vertex holds 2D cartesian coords
type Vertex struct {
	X, Y float64
}

// Scale modifies the vertex by a given scale.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs returns the absolute value of a vertex
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// v := &Vertex{3, 4}
	// fmt.Println(*v, v.Abs())
	// v.Scale(10)
	// fmt.Println(*v, v.Abs())

	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", *v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", *v, v.Abs())
}
