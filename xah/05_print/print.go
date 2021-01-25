package main

import "fmt"

func main() {
	x := "John"
	y := 36

	fmt.Printf("name is %s, age is %d\n", x, y)
	fmt.Println("----")

	// Examples of placeholders

	a := "some value"
	b := 99.9

	type person struct {
		name string
		age  int
	}
	p := person{
		name: "John",
		age:  36,
	}

	fmt.Printf("%%v=%v type=%T\n", a, a)
	fmt.Printf("%%#v=%#v type=%T\n", a, a)
	fmt.Printf("%%q=%q type=%T\n", a, a)
	fmt.Println("----")
	fmt.Printf("%%v=%v type=%T\n", b, b)
	fmt.Printf("%%#v=%#v type=%T\n", b, b)
	fmt.Println("----")
	fmt.Printf("%%v=%v type=%T\n", p, p)
	fmt.Printf("%%#v=%#v type=%T\n", p, p)
	fmt.Printf("%%+v=%+v type=%T\n", p, p)
	fmt.Println("----")

	s := fmt.Sprintf("name is %s, age is %d\n", p.name, p.age)
	fmt.Print(s)
	fmt.Println("----")
}
