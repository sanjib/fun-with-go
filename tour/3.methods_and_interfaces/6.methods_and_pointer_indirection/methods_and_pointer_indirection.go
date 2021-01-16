package main

import "fmt"

// Vertex provides 2D cartesian coords
type Vertex struct {
	X, Y float64
}

// Scale modifies X, Y with given scale.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc modifies X, Y with given scale.
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println("v", v)

	w := Vertex{3, 4}
	ScaleFunc(&w, 10)
	fmt.Println("w", w)

	p := &Vertex{3, 4}
	p.Scale(10)
	fmt.Println("p", *p)

	q := &Vertex{3, 4}
	ScaleFunc(q, 10)
	fmt.Println("q", *q)
}
