package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type personForDec struct {
	Name       string   `json:"fullname"`
	Address    string   `json:"addr"`
	Age        int      `json:"age"`
	FaveColors []string `json:"favecolors"`
}

func main() {
	// json data object
	jsonData := []byte(`
	{
		"fullname": "John Q Public",
		"addr": "987 Main St",
		"age": 45,
		"favecolors": ["Purple", "Yellow", "Gold"]
	}
	`)

	// var p to unmarshal json data
	var p personForDec

	// is json valid
	isValid := json.Valid(jsonData)
	if !isValid {
		log.Println(isValid)
	}

	// unmarshal
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("json decoded as struct:\t%+v\n", p)

	// map
	var m map[string]interface{}
	err = json.Unmarshal(jsonData, &m)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("json decoded as map:\t%v\n", m)
}
