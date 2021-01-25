package main

import "fmt"

func main() {
	fmt.Println("a", len("a"))
	fmt.Println("Ø", len("Ø"))
	fmt.Println("\u2665", len("\u2665"))
	fmt.Println("\U0001f602", len("\U0001f602"))
}
