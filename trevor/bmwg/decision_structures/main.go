package main

import "log"

func main() {
	myVar := "dragon"

	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")
	case "dog":
		log.Println("myVar is set to dog")
	case "fish":
		log.Println("myVar is set to fish")
	default:
		log.Println("myVar is set to default")
	}
}

func main2() {
	b1 := true
	if b1 {
		log.Println("b1:", b1)
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("cat is cat")
	}

	cat = "dog"
	if b1 && cat == "dog" {
		log.Println("cat is dog")
	}
}
