package main

import (
	"fmt"
	"math"
)

// Vertex provides 2D cartesian coords
type Vertex struct {
	X, Y float64
}

// Abs returns the absolute value of a vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// AbsFunc returns the absolute value of a vertex
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
