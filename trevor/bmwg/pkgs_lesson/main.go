package main

import (
	"github.com/sanjib/fake_pkgs_lesson/helpers"
	"log"
)

func main() {
	x1 := helpers.XType{
		TypeName:   "type foo",
		TypeNumber: 111,
	}
	log.Println("x1", x1)
}
