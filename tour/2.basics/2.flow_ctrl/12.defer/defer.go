package main

import "fmt"

func hello() {
	defer fmt.Println("world")
	fmt.Print("Hello ")
}

func main() {
	hello()
}
