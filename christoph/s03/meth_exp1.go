package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name        string
	dateOfBirth time.Time
}

func (p Person) Age() int {
	return int(time.Since(p.dateOfBirth).Hours() / 24 / 365)
}

func (p *Person) ChangeName(s string) {
	p.Name = s
}

func main() {
	p := Person{
		Name:        "Sanjib",
		dateOfBirth: time.Date(1975, 8, 16, 13, 30, 0, 0, &time.Location{}),
	}
	p.ChangeName("Gopher")
	fmt.Println(p.Age())
	fmt.Println(p)
}
