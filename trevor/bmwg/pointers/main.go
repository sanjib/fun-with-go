package main

import "fmt"

func main() {
	var myString string
	myString = "Green"

	changeUsingPointer(&myString)
	fmt.Printf("myString is set to: %q \n", myString)
	fmt.Println("address s:", &myString)
}

func changeUsingPointer(s *string) {
	*s = "Blue"
	fmt.Println("address s:", &s)
}
