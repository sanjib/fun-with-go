package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type personXE struct {
	XMLName    xml.Name `xml:"persondata"`
	FirstName  string   `xml:"first-name"`
	LastName   string   `xml:"last-name"`
	Address    string   `xml:"address"`
	Age        int      `xml:"age,attr"`
	FaveColors []string `xml:"fave-colors"`
}

func main() {
	people := []personXE{
		{FirstName: "Jane", LastName: "Doe", Address: "123 Anywhere St", Age: 35, FaveColors: nil},
		{FirstName: "John", LastName: "Public", Address: "456 Everywhere St", Age: 31, FaveColors: []string{"Purple", "Yellow"}},
	}

	bytes, err := xml.MarshalIndent(people, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Print(xml.Header)
	fmt.Println(string(bytes))
}
