package main

import (
	"log"
	"time"
)

var (
	s1 = "seven"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Sanjib",
		LastName:    "Ahmad",
		PhoneNumber: "123 456 789",
	}
	log.Println("user:", user.FirstName, user.LastName, user.PhoneNumber)
}

func main2() {
	s2 := "six"

	log.Println("s1", s1)
	log.Println("s2", s2)

	log.Println(saySomething("xxx"))
	log.Println(saySomething2())
}

func saySomething(s string) (string, string) {
	return s, "world"
}

func saySomething2() (string, string) {
	return s1, "world"
}
