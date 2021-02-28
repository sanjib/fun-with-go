package main

import "fmt"

type user struct {
	name  string
	email string
}

func createUserV1() user {
	u := user{
		name:  "sanjib",
		email: "sa@sa",
	}
	println("v1 u", &u)
	return u
}

func createUserV2() *user {
	u := user{
		name:  "jibnas",
		email: "ji@ji",
	}
	println("v2 u", &u)
	return &u
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	fmt.Printf("u1 %#v u2 %#v\n", u1, u2)
	println("u1", &u1, "u2", &u2)
}
