package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name       string   `json:"name"`
	Address    string   `json:"address"`
	Age        int      `json:"-"`
	FaveColors []string `json:"favorite_colors,omitempty"`
}

func main() {
	people := []person{
		{"Jane Doe", "123 Anywhere St", 35, nil},
		{"John Public", "456 Everywhere St", 31, []string{"Purple", "Yellow"}},
	}
	bytes, err := json.Marshal(people)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("json encoded:\t", string(bytes))

	bytes2, err := json.MarshalIndent(people, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("json encoded:\t", string(bytes2))
}
