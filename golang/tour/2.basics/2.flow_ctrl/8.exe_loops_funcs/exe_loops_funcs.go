package main

import (
	"fmt"
	"math"
)

// Sqrt calculates the square root of the given number.
func Sqrt(x float64) float64 {
	z := 1.0
	prev2, prev1 := 0.0, 0.0
	i := 1
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if z == prev1 || z == prev2 {
			fmt.Printf("took %d iterations\n", i)
			return z
		}
		prev2 = prev1
		prev1 = z
		i++
	}
}

func main() {
	x := 923457
	fmt.Println("cust Sqrt", Sqrt(float64(x)))
	fmt.Println("math.Sqrt", math.Sqrt(float64(x)))
}
