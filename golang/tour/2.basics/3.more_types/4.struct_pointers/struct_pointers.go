package main

import "fmt"

// Vertex holds cartesian coordinates
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 22
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Println(v)
}
