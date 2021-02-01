package main

import (
	"fmt"
	"math"
)

// Vertex holds 2D cartesian coordinates.
type Vertex struct {
	X, Y float64
}

// Abs returns the absolute value of a vertex.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale modified the X, Y coords of a vertex to scale.
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleWithoutPointer returns a modified vertex with X, Y modifed to scale.
func ScaleWithoutPointer(v Vertex, f float64) Vertex {
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	//Scale(&v, 10)
	v = ScaleWithoutPointer(v, 10)
	fmt.Println(v)
	fmt.Println(Abs(v))
}
