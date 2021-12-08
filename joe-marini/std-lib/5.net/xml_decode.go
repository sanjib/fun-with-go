package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type personXD struct {
	XMLName    xml.Name `xml:"persondata"`
	FirstName  string   `xml:"first-name"`
	LastName   string   `xml:"last-name"`
	Address    string   `xml:"address"`
	Age        int      `xml:"age,attr"`
	FaveColors []string `xml:"fave-colors"`
}

func main() {
	xmldata := `
	<persondata age="31">
	  <first-name>John</first-name>
	  <last-name>Public</last-name>
	  <address>456 Everywhere St</address>
	  <fave-colors>Purple</fave-colors>
	  <fave-colors>Yellow</fave-colors>
	</persondata>
	`
	var p personXD
	err := xml.Unmarshal([]byte(xmldata), &p)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%#v\n", p)
}
