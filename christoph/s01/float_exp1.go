package main

import "fmt"

func main() {
	i := -1.0
	res := i * 0
	fmt.Println(res)

	f1 := 1.1
	f2 := 0.1
	fmt.Println(f1 + f2)

	var f64 float64 = 642545754
	var f32 float32 = 642545754
	fmt.Printf("%.0f\n", f64)
	fmt.Printf("%.0f\n", f32)
	fmt.Println(f64 - float64(f32))

	var f32a float32 = 999956000
	fmt.Printf("f32a %.0f\n", f32a)
}
