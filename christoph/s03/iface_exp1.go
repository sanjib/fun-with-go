package main

import "fmt"

type Helloer interface {
	Hello(string)
}

type Greeting string

func (g Greeting) Hello(s string) {
	fmt.Println(g, s)
}

type Invitation struct {
	event string
}

func (i *Invitation) Hello(s string) {
	fmt.Println(i.event, s)
}

func main() {
	var h Helloer

	h = Greeting("H.e.l.l.o.")
	h.Hello("sanjib")

	fmt.Printf("%#v\n", h)

	h = &Invitation{event: "H.E.L.L.O."}
	h.Hello("jibnas")

	fmt.Printf("%#v\n", h)

	i := Invitation{event: "Howdy"}
	i.Hello("san")
	fmt.Printf("%#v\n", i)
}
