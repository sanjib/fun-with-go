package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	cu := location{-4.61, 137.41}
	bytes, err := json.Marshal(cu)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v", string(bytes))
}
