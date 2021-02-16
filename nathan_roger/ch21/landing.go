package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type location struct {
		Name string  `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	locations := []location{
		{"Bradbury Landing", -4, 137},
		{"Columbia Memorial", -14, 174},
		{"Challenger Memorial", -1, 354},
	}
	bytes, err := json.MarshalIndent(locations, "", "\t")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%v", string(bytes))

}
