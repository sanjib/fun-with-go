package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
  {
    "first_name": "Clark",
    "last_name": "Kent",
    "hair_color": "black",
    "has_dog": true
  },
  {
    "first_name": "Bruce",
    "last_name": "Wayne",
    "hair_color": "black",
    "has_dog": false
  }
]`

	// read JSON
	var persons []Person
	err := json.Unmarshal([]byte(myJson), &persons)
	if err != nil {
		log.Println("error unmarshalling json", err)
	}
	log.Printf("-> unmarshalled: %v", persons)

	// write JSON
	p1 := Person{
		FirstName: "Wally",
		LastName:  "West",
		HairColor: "red",
		HasDog:    false,
	}
	p2 := Person{
		FirstName: "Diana",
		LastName:  "Prince",
		HairColor: "black",
		HasDog:    false,
	}

	jsonPeople, err := json.MarshalIndent([]Person{p1, p2}, "", "  ")
	if err != nil {
		log.Println(err)
	}
	log.Println("-> json:", string(jsonPeople))

}
