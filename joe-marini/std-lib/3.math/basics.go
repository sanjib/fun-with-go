package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.Ceil(math.Pi))
	fmt.Println(math.Floor(math.Pi))
	fmt.Printf("%.2f\n", math.Trunc(math.Pi))
	fmt.Printf("%.2f\n", math.Floor(math.Pi))
	fmt.Println(math.Max(math.Pi, math.Ln2))
	fmt.Println(math.Min(math.Pi, math.Ln2))
	fmt.Println(17 % 5)
	fmt.Println(math.Mod(17.0, 5.0))
	fmt.Printf("%.1f\n", math.Round(10.5))
	fmt.Printf("%.1f\n", math.Round(-10.5))
	fmt.Printf("%.1f\n", math.RoundToEven(10.5))
	fmt.Printf("%.1f\n", math.RoundToEven(11.5))
}
