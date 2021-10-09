package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	users := []User{
		{"John", "Smith", "john@smith.com", 30},
		{"Mary", "Jone", "mary@jones.com", 45},
		{"Sally", "Brown", "sally@brown.com", 17},
	}

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}

	var m1 map[string]string
	log.Println("m1", m1)

	m2 := make(map[string]string)
	log.Println("m2", m2)

}

func main3() {
	animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"
	for animalType, animalName := range animals {
		log.Println(animalType, animalName)
	}

	firstLine := "Once upon a midnight dreary 😁"
	for i, val := range firstLine {
		log.Println(i, val)
	}
	log.Println("len firstLine", len(firstLine))
	log.Println("address firstLine", &firstLine)
	firstLine = "x"
	log.Println("address firstLine", &firstLine)

}

func main2() {
	for i := 0; i < 3; i++ {
		log.Println(i)
	}
	log.Println("--")

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}
	log.Println("--")

}
