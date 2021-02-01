package main

import "fmt"

// Vertex contains cartesian coordinates
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
