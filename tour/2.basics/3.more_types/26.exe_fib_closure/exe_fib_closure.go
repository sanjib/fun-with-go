package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := [2]int{}
	return func() int {
		sum := a[0] + a[1]
		a[0] = a[1]
		a[1] = sum
		if a[0] == 0 && a[1] == 0 {
			a[0] = 1
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
