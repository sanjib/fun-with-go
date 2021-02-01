package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", f, f)
	fmt.Printf("%v %T\n", z, z)
}
