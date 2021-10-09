package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	myVar := myStruct{
		FirstName: "John",
	}
	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar", myVar.printFirstName())
	log.Println("myVar2", myVar2.printFirstName())
}
