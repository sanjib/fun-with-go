package main

import "fmt"

func main() {
	//d := dog{}
	//speak(d)
	//
	//c := cat{}
	//speak(c)
	//
	//ct := cutie{}
	//speak(ct)

	subjects := []talk{
		dog{},
		cat{},
		cutie{},
	}
	for _, subject := range subjects {
		speak(subject)
	}
}

func speak(t talk) {
	fmt.Println(t.talk())
}

// interface is a proxy for a real type of what methods it can perform
type talk interface {
	talk() string
}

// dog is of type struct
type dog struct{}

func (d dog) talk() string {
	return "woof woof"
}

// cat is of type struct
type cat struct{}

func (c cat) talk() string {
	return "meow"
}

// cutie is of type struct
type cutie struct {
	cat
}
