package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10.0
	fmt.Println("abs", math.Abs(x), math.Abs(-x))
	fmt.Println("pow", math.Pow(x, 2.0))
	fmt.Println("exp", math.E, math.Exp(x))
	fmt.Println("cos", math.Cos(math.Pi))
	fmt.Println("sin", math.Sin(math.Pi/2))
	fmt.Println("tan", math.Tan(math.Pi/2))
	fmt.Println("log", math.Log(10))
	fmt.Println("sqrt", math.Sqrt(25))
	fmt.Println("cbrt", math.Cbrt(125))
	fmt.Println("hypot", math.Hypot(30, 40))

}
