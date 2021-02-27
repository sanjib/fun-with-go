package main

import "fmt"

func trace(name string) func() {
	// 1. Print "Entering <name>"
	fmt.Println("Entering", name)
	// 2. return a func() that prints "Leaving <name>"
	return func() {
		fmt.Println("Leaving", name)
	}
}
func f() {
	// this defers execution of both trace and the return function
	defer func() {
		f := trace("f")
		f()
	}()

	// this immediately executes trace and defers execution of the returned function
	//defer trace("f")()

	fmt.Println("Doing something")
}
func main() {
	fmt.Println("Before f")
	f()
	fmt.Println("After f")
}
