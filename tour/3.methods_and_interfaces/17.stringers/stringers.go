package main

import "fmt"

// Person is a struct
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"John Doe", 42}
	b := Person{"Jane Doe", 40}

	fmt.Println(a, b)
}
