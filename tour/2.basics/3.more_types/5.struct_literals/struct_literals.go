package main

import "fmt"

// Vertex holds cartesian coordinates
type Vertex struct {
	X, Y int
}

func main() {
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	p := &Vertex{1, 2}

	fmt.Println(v1, v2, v3, p)
}
